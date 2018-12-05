package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	meteringv1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	meteringutil "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (f *Framework) CreateMeteringReport(report *meteringv1alpha1.Report) error {
	_, err := f.MeteringClient.Reports(f.Namespace).Create(report)
	return err
}

func (f *Framework) CreateMeteringScheduledReport(report *meteringv1alpha1.ScheduledReport) error {
	_, err := f.MeteringClient.ScheduledReports(f.Namespace).Create(report)
	return err
}

func (f *Framework) GetMeteringReport(name string) (*meteringv1alpha1.Report, error) {
	return f.MeteringClient.Reports(f.Namespace).Get(name, meta.GetOptions{})
}

func (f *Framework) GetMeteringScheduledReport(name string) (*meteringv1alpha1.ScheduledReport, error) {
	return f.MeteringClient.ScheduledReports(f.Namespace).Get(name, meta.GetOptions{})
}

func (f *Framework) NewSimpleReport(name, queryName string, reportingStart, reportingEnd *time.Time) *meteringv1alpha1.Report {
	var start, end *meta.Time
	if reportingStart != nil {
		start = &meta.Time{*reportingStart}
	}
	if reportingEnd != nil {
		end = &meta.Time{*reportingEnd}
	}
	return &meteringv1alpha1.Report{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: meteringv1alpha1.ReportSpec{
			GenerationQueryName: queryName,
			RunImmediately:      true,
			ReportingStart:      start,
			ReportingEnd:        end,
		},
	}
}

func (f *Framework) NewSimpleScheduledReport(name, queryName string, schedule *meteringv1alpha1.ScheduledReportSchedule, reportingStart, reportingEnd *time.Time) *meteringv1alpha1.ScheduledReport {
	var start, end *meta.Time
	if reportingStart != nil {
		start = &meta.Time{*reportingStart}
	}
	if reportingEnd != nil {
		end = &meta.Time{*reportingEnd}
	}
	return &meteringv1alpha1.ScheduledReport{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: meteringv1alpha1.ScheduledReportSpec{
			GenerationQueryName: queryName,
			Schedule:            schedule,
			ReportingStart:      start,
			ReportingEnd:        end,
		},
	}
}

func (f *Framework) RequireScheduledReportSuccessfullyRuns(t *testing.T, report *meteringv1alpha1.ScheduledReport, waitTimeout time.Duration) {
	err := f.MeteringClient.ScheduledReports(f.Namespace).Delete(report.Name, nil)
	assert.Condition(t, func() bool {
		return err == nil || errors.IsNotFound(err)
	}, "failed to ensure scheduled report doesn't exist before creating scheduled report")

	t.Logf("creating scheduled report %s", report.Name)
	err = f.CreateMeteringScheduledReport(report)
	require.NoError(t, err, "creating scheduled report should succeed")

	reportEnd := report.Spec.ReportingEnd.Time

	err = wait.PollImmediate(time.Second*5, waitTimeout, func() (bool, error) {
		// poll the status
		newReport, err := f.GetMeteringScheduledReport(report.Name)
		if err != nil {
			return false, err
		}
		cond := meteringutil.GetScheduledReportCondition(newReport.Status, meteringv1alpha1.ScheduledReportFailure)
		if cond != nil && cond.Status == v1.ConditionTrue {
			return false, fmt.Errorf("report is failed, reason: %s, message: %s", cond.Reason, cond.Message)
		}

		if newReport.Status.TableName == "" {
			t.Logf("ScheduledReport %s table isn't created yet, status: %+v", report.Name, newReport.Status)
			return false, nil
		}

		// If the last reportTime is updated, that means this report
		// has been run at least once.
		if newReport.Status.LastReportTime == nil {
			t.Logf("report LastReportTime is unset")
			return false, nil
		}
		t.Logf("report LastReportTime: %s", newReport.Status.LastReportTime)
		if !(newReport.Status.LastReportTime.Time.After(reportEnd) || newReport.Status.LastReportTime.Time.Equal(reportEnd)) {
			t.Logf("LastReportTime %s newReport.Status.LastReportTime is not >= reportEnd %s", newReport.Status.LastReportTime, reportEnd)
			return false, nil
		}
		return true, nil
	})
	require.NoError(t, err, "expected getting ScheduledReport to not timeout")

}

func (f *Framework) GetScheduledReportResults(t *testing.T, report *meteringv1alpha1.ScheduledReport, waitTimeout time.Duration) []map[string]interface{} {
	var reportResults []map[string]interface{}
	var reportData []byte

	queryParams := map[string]string{
		"name":   report.Name,
		"format": "json",
	}
	err := wait.PollImmediate(time.Second*5, waitTimeout, func() (bool, error) {
		req := f.NewReportingOperatorSVCRequest("/api/v1/scheduledreports/get", queryParams)
		result := req.Do()
		resp, err := result.Raw()
		require.NoError(t, err, "fetching ScheduledReport results should be successful")

		var statusCode int
		result.StatusCode(&statusCode)
		if statusCode == http.StatusAccepted {
			t.Logf("report is still running")
			return false, nil
		}

		require.Equal(t, http.StatusOK, statusCode, "http response status code should be ok")
		err = json.Unmarshal(resp, &reportResults)
		require.NoError(t, err, "expected to unmarshal response")
		reportData = resp
		if len(reportResults) == 0 {
			t.Logf("ScheduledReport %s has 0 results", report.Name)
			return false, nil
		}
		return true, nil
	})
	require.NoError(t, err, "expected ScheduledReport to have 1 row of results before timing out")
	assert.NotEmpty(t, reportResults, "reports should return at least 1 row")

	***REMOVED***leName := path.Join(f.ReportOutputDirectory, fmt.Sprintf("%s-scheduled-report.json", report.Name))
	err = ioutil.WriteFile(***REMOVED***leName, reportData, os.ModePerm)
	require.NoError(t, err, "expected writing report results to disk not to error")
	return reportResults
}
