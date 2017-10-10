/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this ***REMOVED***le except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the speci***REMOVED***c language governing permissions and
limitations under the License.
*/

package generators

import (
	"fmt"
	"path/***REMOVED***lepath"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"

	"k8s.io/code-generator/cmd/client-gen/generators/util"
	clientgentypes "k8s.io/code-generator/cmd/client-gen/types"

	"github.com/golang/glog"
)

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	pluralExceptions := map[string]string{
		"Endpoints": "Endpoints",
	}
	return namer.NameSystems{
		"public":             namer.NewPublicNamer(0),
		"private":            namer.NewPrivateNamer(0),
		"raw":                namer.NewRawNamer("", nil),
		"publicPlural":       namer.NewPublicPluralNamer(pluralExceptions),
		"allLowercasePlural": namer.NewAllLowercasePluralNamer(pluralExceptions),
		"lowercaseSingular":  &lowercaseSingularNamer{},
	}
}

// lowercaseSingularNamer implements Namer
type lowercaseSingularNamer struct{}

// Name returns t's name in all lowercase.
func (n *lowercaseSingularNamer) Name(t *types.Type) string {
	return strings.ToLower(t.Name.Name)
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

// generatedBy returns information about the arguments used to invoke
// lister-gen.
func generatedBy() string {
	return fmt.Sprintf("\n// This ***REMOVED***le was automatically generated by informer-gen\n\n")
}

// objectMetaForPackage returns the type of ObjectMeta used by package p.
func objectMetaForPackage(p *types.Package) (*types.Type, bool, error) {
	generatingForPackage := false
	for _, t := range p.Types {
		if !util.MustParseClientGenTags(t.SecondClosestCommentLines).GenerateClient {
			continue
		}
		generatingForPackage = true
		for _, member := range t.Members {
			if member.Name == "ObjectMeta" {
				return member.Type, isInternal(member), nil
			}
		}
	}
	if generatingForPackage {
		return nil, false, fmt.Errorf("unable to ***REMOVED***nd ObjectMeta for any types in package %s", p.Path)
	}
	return nil, false, nil
}

// isInternal returns true if the tags for a member do not contain a json tag
func isInternal(m types.Member) bool {
	return !strings.Contains(m.Tags, "json")
}

func packageForGroup(base string, group clientgentypes.Group) string {
	return ***REMOVED***lepath.Join(base, group.NonEmpty())
}

func packageForInternalInterfaces(base string) string {
	return ***REMOVED***lepath.Join(base, "internalinterfaces")
}

func vendorless(p string) string {
	if pos := strings.LastIndex(p, "/vendor/"); pos != -1 {
		return p[pos+len("/vendor/"):]
	}
	return p
}

// Packages makes the client package de***REMOVED***nition.
func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		glog.Fatalf("Failed loading boilerplate: %v", err)
	}

	boilerplate = append(boilerplate, []byte(generatedBy())...)

	customArgs, ok := arguments.CustomArgs.(*CustomArgs)
	if !ok {
		glog.Fatalf("Wrong CustomArgs type: %T", arguments.CustomArgs)
	}

	internalVersionPackagePath := ***REMOVED***lepath.Join(arguments.OutputPackagePath)
	externalVersionPackagePath := ***REMOVED***lepath.Join(arguments.OutputPackagePath)
	if !customArgs.SingleDirectory {
		internalVersionPackagePath = ***REMOVED***lepath.Join(arguments.OutputPackagePath, "internalversion")
		externalVersionPackagePath = ***REMOVED***lepath.Join(arguments.OutputPackagePath, "externalversions")
	}

	var packageList generator.Packages
	typesForGroupVersion := make(map[clientgentypes.GroupVersion][]*types.Type)

	externalGroupVersions := make(map[string]clientgentypes.GroupVersions)
	internalGroupVersions := make(map[string]clientgentypes.GroupVersions)
	for _, inputDir := range arguments.InputDirs {
		p := context.Universe.Package(vendorless(inputDir))

		objectMeta, internal, err := objectMetaForPackage(p)
		if err != nil {
			glog.Fatal(err)
		}
		if objectMeta == nil {
			// no types in this package had genclient
			continue
		}

		var gv clientgentypes.GroupVersion
		var targetGroupVersions map[string]clientgentypes.GroupVersions

		if internal {
			lastSlash := strings.LastIndex(p.Path, "/")
			if lastSlash == -1 {
				glog.Fatalf("error constructing internal group version for package %q", p.Path)
			}
			gv.Group = clientgentypes.Group(p.Path[lastSlash+1:])
			targetGroupVersions = internalGroupVersions
		} ***REMOVED*** {
			parts := strings.Split(p.Path, "/")
			gv.Group = clientgentypes.Group(parts[len(parts)-2])
			gv.Version = clientgentypes.Version(parts[len(parts)-1])
			targetGroupVersions = externalGroupVersions
		}

		// If there's a comment of the form "// +groupName=somegroup" or
		// "// +groupName=somegroup.foo.bar.io", use the ***REMOVED***rst ***REMOVED***eld (somegroup) as the name of the
		// group when generating.
		if override := types.ExtractCommentTags("+", p.DocComments)["groupName"]; override != nil {
			gv.Group = clientgentypes.Group(strings.SplitN(override[0], ".", 2)[0])
		}

		var typesToGenerate []*types.Type
		for _, t := range p.Types {
			tags := util.MustParseClientGenTags(t.SecondClosestCommentLines)
			if !tags.GenerateClient || tags.NoVerbs || !tags.HasVerb("list") || !tags.HasVerb("watch") {
				continue
			}

			typesToGenerate = append(typesToGenerate, t)

			if _, ok := typesForGroupVersion[gv]; !ok {
				typesForGroupVersion[gv] = []*types.Type{}
			}
			typesForGroupVersion[gv] = append(typesForGroupVersion[gv], t)
		}
		if len(typesToGenerate) == 0 {
			continue
		}

		icGroupName := namer.IC(gv.Group.NonEmpty())
		groupVersionsEntry, ok := targetGroupVersions[icGroupName]
		if !ok {
			groupVersionsEntry = clientgentypes.GroupVersions{
				Group: gv.Group,
			}
		}
		groupVersionsEntry.Versions = append(groupVersionsEntry.Versions, gv.Version)
		targetGroupVersions[icGroupName] = groupVersionsEntry

		orderer := namer.Orderer{Namer: namer.NewPrivateNamer(0)}
		typesToGenerate = orderer.OrderTypes(typesToGenerate)

		if internal {
			packageList = append(packageList, versionPackage(internalVersionPackagePath, gv, boilerplate, typesToGenerate, customArgs.InternalClientSetPackage, customArgs.ListersPackage))
		} ***REMOVED*** {
			packageList = append(packageList, versionPackage(externalVersionPackagePath, gv, boilerplate, typesToGenerate, customArgs.VersionedClientSetPackage, customArgs.ListersPackage))
		}
	}

	if len(externalGroupVersions) != 0 {
		packageList = append(packageList, factoryInterfacePackage(externalVersionPackagePath, boilerplate, customArgs.VersionedClientSetPackage, typesForGroupVersion))
		packageList = append(packageList, factoryPackage(externalVersionPackagePath, boilerplate, externalGroupVersions, customArgs.VersionedClientSetPackage, typesForGroupVersion))
		for _, groupVersionsEntry := range externalGroupVersions {
			packageList = append(packageList, groupPackage(externalVersionPackagePath, groupVersionsEntry, boilerplate))
		}
	}

	if len(internalGroupVersions) != 0 {
		packageList = append(packageList, factoryInterfacePackage(internalVersionPackagePath, boilerplate, customArgs.InternalClientSetPackage, typesForGroupVersion))
		packageList = append(packageList, factoryPackage(internalVersionPackagePath, boilerplate, internalGroupVersions, customArgs.InternalClientSetPackage, typesForGroupVersion))
		for _, groupVersionsEntry := range internalGroupVersions {
			packageList = append(packageList, groupPackage(internalVersionPackagePath, groupVersionsEntry, boilerplate))
		}
	}

	return packageList
}

func isInternalVersion(gv clientgentypes.GroupVersion) bool {
	return len(gv.Version) == 0
}

func factoryPackage(basePackage string, boilerplate []byte, groupVersions map[string]clientgentypes.GroupVersions, clientSetPackage string, typesForGroupVersion map[clientgentypes.GroupVersion][]*types.Type) generator.Package {
	return &generator.DefaultPackage{
		PackageName: ***REMOVED***lepath.Base(basePackage),
		PackagePath: basePackage,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &factoryGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "factory",
				},
				outputPackage:             basePackage,
				imports:                   generator.NewImportTracker(),
				groupVersions:             groupVersions,
				clientSetPackage:          clientSetPackage,
				internalInterfacesPackage: packageForInternalInterfaces(basePackage),
			})

			generators = append(generators, &genericGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "generic",
				},
				outputPackage:        basePackage,
				imports:              generator.NewImportTracker(),
				groupVersions:        groupVersions,
				typesForGroupVersion: typesForGroupVersion,
			})

			return generators
		},
	}
}

func factoryInterfacePackage(basePackage string, boilerplate []byte, clientSetPackage string, typesForGroupVersion map[clientgentypes.GroupVersion][]*types.Type) generator.Package {
	packagePath := packageForInternalInterfaces(basePackage)

	return &generator.DefaultPackage{
		PackageName: ***REMOVED***lepath.Base(packagePath),
		PackagePath: packagePath,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &factoryInterfaceGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "factory_interfaces",
				},
				outputPackage:    packagePath,
				imports:          generator.NewImportTracker(),
				clientSetPackage: clientSetPackage,
			})

			return generators
		},
	}
}

func groupPackage(basePackage string, groupVersions clientgentypes.GroupVersions, boilerplate []byte) generator.Package {
	packagePath := ***REMOVED***lepath.Join(basePackage, strings.ToLower(groupVersions.Group.NonEmpty()))

	return &generator.DefaultPackage{
		PackageName: strings.ToLower(groupVersions.Group.NonEmpty()),
		PackagePath: packagePath,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &groupInterfaceGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "interface",
				},
				outputPackage:             packagePath,
				groupVersions:             groupVersions,
				imports:                   generator.NewImportTracker(),
				internalInterfacesPackage: packageForInternalInterfaces(basePackage),
			})
			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			tags := util.MustParseClientGenTags(t.SecondClosestCommentLines)
			return tags.GenerateClient && tags.HasVerb("list") && tags.HasVerb("watch")
		},
	}
}

func versionPackage(basePackage string, gv clientgentypes.GroupVersion, boilerplate []byte, typesToGenerate []*types.Type, clientSetPackage, listersPackage string) generator.Package {
	packagePath := ***REMOVED***lepath.Join(basePackage, strings.ToLower(gv.Group.NonEmpty()), strings.ToLower(gv.Version.NonEmpty()))

	return &generator.DefaultPackage{
		PackageName: strings.ToLower(gv.Version.NonEmpty()),
		PackagePath: packagePath,
		HeaderText:  boilerplate,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = append(generators, &versionInterfaceGenerator{
				DefaultGen: generator.DefaultGen{
					OptionalName: "interface",
				},
				outputPackage: packagePath,
				imports:       generator.NewImportTracker(),
				types:         typesToGenerate,
				internalInterfacesPackage: packageForInternalInterfaces(basePackage),
			})

			for _, t := range typesToGenerate {
				generators = append(generators, &informerGenerator{
					DefaultGen: generator.DefaultGen{
						OptionalName: strings.ToLower(t.Name.Name),
					},
					outputPackage:             packagePath,
					groupVersion:              gv,
					typeToGenerate:            t,
					imports:                   generator.NewImportTracker(),
					clientSetPackage:          clientSetPackage,
					listersPackage:            listersPackage,
					internalInterfacesPackage: packageForInternalInterfaces(basePackage),
				})
			}
			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			tags := util.MustParseClientGenTags(t.SecondClosestCommentLines)
			return tags.GenerateClient && tags.HasVerb("list") && tags.HasVerb("watch")
		},
	}
}
