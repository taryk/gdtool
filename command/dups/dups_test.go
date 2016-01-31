package dups

import (
	"testing"
	"os"
	"flag"
	"github.com/fatih/color"

	. "github.com/taryk/gdtool/core"
)

type testpair struct {
	input FileList
	grouped FileTreeMap
	duplicates []*FileDetailsMap
}

var (
	empty = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF0",
		Name: "Empty",
		Path: "/",
		Size: 12345,
		Md5sum: "3052c318de21276469bcf3af4654900a",
		IsDir: true,
	}
	vesuvious = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF1",
		Name: "Vesuvius.jpg",
		Path: "/",
		Size: 12345,
		Md5sum: "497e262d5f8f93e9e13b271b9e0c461f",
		IsDir: false,
	}
	solfatara = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQM21",
		Name: "Solfatara.jpg",
		Path: "/",
		Size: 12346,
		Md5sum: "8da7cd386547fc7a82deaeb70e80e58f",
		IsDir: false,
	}
	nuovo = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQM22",
		Name: "Nuovo.jpg",
		Path: "/",
		Size: 12346,
		Md5sum: "cca34fc9978c9d1123ae903d6e17d8f2",
		IsDir: false,
	}
	averna = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQM23",
		Name: "Averna.jpg",
		Path: "/",
		Size: 12346,
		Md5sum: "78466284f3bffec29c043683cc281401",
		IsDir: false,
	}
	etna = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF3",
		Name: "Etna.jpg",
		Path: "/",
		Size: 12347,
		Md5sum: "7d35fb277155363fac749b791be06fe0",
		IsDir: false,
	}
	volcano = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF4",
		Name: "Volcano.jpg",
		Path: "/",
		Size: 12348,
		Md5sum: "3516b45c828847be506a34bda1727761",
		IsDir: false,
	}
	stromboli = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF5",
		Name: "Stromboli.jpg",
		Path: "/",
		Size: 12349,
		Md5sum: "505f07f2b50446f955e90f26c92b81a4",
		IsDir: false,
	}
	stromboliDifferentMd5sum = &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQME5",
		Name: "Stromboli.jpg",
		Path: "/",
		Size: 12349,
		Md5sum: "ac446dc11784fdf0cb7b4e83eb41775c",
		IsDir: false,
	}
	campiflegrei = campiflegreiDir(solfatara, nuovo, averna)
	italianMainlandVolcanoes = italianMainlandVolcanoesDir(
		vesuvious, solfatara, nuovo, averna,
	)
	partialItalianMainlandVolcanoes = partialItalianMainlandVolcanoesDir(
		vesuvious,
	)
	sicilianVolcanoes = sicilianVolcanoesDir(etna)
	partialSicilianVolcanoes = partialSicilianVolcanoesDir()
	volcanicIslands = volcanicIslandsDir(volcano, stromboli)
	partialVolcanicIslands = partialVolcanicIslandsDir(volcano)
	italianVolcanoes = italianVolcanoesDir(
		italianMainlandVolcanoes,
		sicilianVolcanoes,
		volcanicIslands,
	)
	partialItalianVolcanoes = partialItalianVolcanoesDir(
		partialItalianMainlandVolcanoes,
		partialSicilianVolcanoes,
		partialVolcanicIslands,
	)
)

func campiflegreiDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQM20",
		Name: "Campi Flegrei",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}

func italianMainlandVolcanoesDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF6",
		Name: "Italian Mainland Volcanoes",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}

func partialItalianMainlandVolcanoesDir(children ...*FileDetails) *FileDetails{
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF6",
		Name: "Italian Mainland Volcanoes",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}
func sicilianVolcanoesDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF7",
		Name: "Sicilian Volcanoes",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}
func partialSicilianVolcanoesDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF7",
		Name: "Sicilian Volcanoes",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}
func volcanicIslandsDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF8",
		Name: "Volcanic Islands",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}
func partialVolcanicIslandsDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF8",
		Name: "Volcanic Islands",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}
func italianVolcanoesDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF9",
		Name: "Italian Volcanoes",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}
func partialItalianVolcanoesDir(children ...*FileDetails) *FileDetails {
	return &FileDetails{
		Id: "0B--_5Ov7G9k11GtGU3ZTSzVQMF9",
		Name: "Italian Volcanoes",
		Path: "/",
		Size: 0,
		IsDir: true,
		Children: children,
	}
}

var tests = []testpair{
	{
		input: FileList{
			italianMainlandVolcanoesDir(
				vesuvious,
				campiflegreiDir(
					solfatara,
					nuovo,
					averna,
				),
			),
			sicilianVolcanoesDir(etna),
			volcanicIslandsDir(stromboli, volcano),
		},
		grouped: FileTreeMap{
			italianMainlandVolcanoes.Name: []*FileDetailsMap{
				&FileDetailsMap{
					Id: italianMainlandVolcanoes.Id,
					Name: italianMainlandVolcanoes.Name,
					Path: italianMainlandVolcanoes.Path,
					Size: italianMainlandVolcanoes.Size,
					Md5sum: italianMainlandVolcanoes.Md5sum,
					IsDir: italianMainlandVolcanoes.IsDir,
					Children: FileTreeMap{
						vesuvious.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: vesuvious.Id,
								Name: vesuvious.Name,
								Path: vesuvious.Path,
								Size: vesuvious.Size,
								Md5sum: vesuvious.Md5sum,
								IsDir: vesuvious.IsDir,
							},
						},
						campiflegrei.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: campiflegrei.Id,
								Name: campiflegrei.Name,
								Path: campiflegrei.Path,
								Size: campiflegrei.Size,
								Md5sum: campiflegrei.Md5sum,
								IsDir: campiflegrei.IsDir,
								Children: FileTreeMap{
									solfatara.Name: []*FileDetailsMap{
										&FileDetailsMap{
											Id: solfatara.Id,
											Name: solfatara.Name,
											Path: solfatara.Path,
											Size: solfatara.Size,
											Md5sum: solfatara.Md5sum,
											IsDir: solfatara.IsDir,
										},
									},
									nuovo.Name: []*FileDetailsMap{
										&FileDetailsMap{
											Id: nuovo.Id,
											Name: nuovo.Name,
											Path: nuovo.Path,
											Size: nuovo.Size,
											Md5sum: nuovo.Md5sum,
											IsDir: nuovo.IsDir,
										},
									},
									averna.Name: []*FileDetailsMap{
										&FileDetailsMap{
											Id: averna.Id,
											Name: averna.Name,
											Path: averna.Path,
											Size: averna.Size,
											Md5sum: averna.Md5sum,
											IsDir: averna.IsDir,
										},
									},
								},
							},
						},
					},
				},
			},
			sicilianVolcanoes.Name: []*FileDetailsMap{
				&FileDetailsMap{
					Id: sicilianVolcanoes.Id,
					Name: sicilianVolcanoes.Name,
					Path: sicilianVolcanoes.Path,
					Size: sicilianVolcanoes.Size,
					Md5sum: sicilianVolcanoes.Md5sum,
					IsDir: sicilianVolcanoes.IsDir,
					Children: FileTreeMap{
						etna.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: etna.Id,
								Name: etna.Name,
								Path: etna.Path,
								Size: etna.Size,
								Md5sum: etna.Md5sum,
								IsDir: etna.IsDir,
							},
						},
					},
				},
			},
			volcanicIslands.Name: []*FileDetailsMap{
				&FileDetailsMap{
					Id: volcanicIslands.Id,
					Name: volcanicIslands.Name,
					Path: volcanicIslands.Path,
					Size: volcanicIslands.Size,
					Md5sum: volcanicIslands.Md5sum,
					IsDir: volcanicIslands.IsDir,
					Children: FileTreeMap{
						stromboli.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: stromboli.Id,
								Name: stromboli.Name,
								Path: stromboli.Path,
								Size: stromboli.Size,
								Md5sum: stromboli.Md5sum,
								IsDir: stromboli.IsDir,
							},
						},
						volcano.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: volcano.Id,
								Name: volcano.Name,
								Path: volcano.Path,
								Size: volcano.Size,
								Md5sum: volcano.Md5sum,
								IsDir: volcano.IsDir,
							},
						},
					},
				},
			},
		},
		duplicates: []*FileDetailsMap{},
	},
	{
		input: FileList{
			italianMainlandVolcanoesDir(
				vesuvious,
				vesuvious,
				campiflegreiDir(
					solfatara,
					solfatara,
					solfatara,
					nuovo,
					averna,
					averna,
				),
			),
			sicilianVolcanoesDir(etna),
			sicilianVolcanoesDir(etna),
			volcanicIslandsDir(stromboli, volcano),
			volcanicIslandsDir(stromboli, volcano),
			volcanicIslandsDir(stromboli, volcano),
		},
		grouped: FileTreeMap{
			italianMainlandVolcanoes.Name: []*FileDetailsMap{
				&FileDetailsMap{
					Id: italianMainlandVolcanoes.Id,
					Name: italianMainlandVolcanoes.Name,
					Path: italianMainlandVolcanoes.Path,
					Size: italianMainlandVolcanoes.Size,
					Md5sum: italianMainlandVolcanoes.Md5sum,
					IsDir: italianMainlandVolcanoes.IsDir,
					Children: FileTreeMap{
						vesuvious.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: vesuvious.Id,
								Name: vesuvious.Name,
								Path: vesuvious.Path,
								Size: vesuvious.Size,
								Md5sum: vesuvious.Md5sum,
								IsDir: vesuvious.IsDir,
							},
							&FileDetailsMap{
								Id: vesuvious.Id,
								Name: vesuvious.Name,
								Path: vesuvious.Path,
								Size: vesuvious.Size,
								Md5sum: vesuvious.Md5sum,
								IsDir: vesuvious.IsDir,
							},
						},
						campiflegrei.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: campiflegrei.Id,
								Name: campiflegrei.Name,
								Path: campiflegrei.Path,
								Size: campiflegrei.Size,
								Md5sum: campiflegrei.Md5sum,
								IsDir: campiflegrei.IsDir,
								Children: FileTreeMap{
									solfatara.Name: []*FileDetailsMap{
										&FileDetailsMap{
											Id: solfatara.Id,
											Name: solfatara.Name,
											Path: solfatara.Path,
											Size: solfatara.Size,
											Md5sum: solfatara.Md5sum,
											IsDir: solfatara.IsDir,
										},
										&FileDetailsMap{
											Id: solfatara.Id,
											Name: solfatara.Name,
											Path: solfatara.Path,
											Size: solfatara.Size,
											Md5sum: solfatara.Md5sum,
											IsDir: solfatara.IsDir,
										},
										&FileDetailsMap{
											Id: solfatara.Id,
											Name: solfatara.Name,
											Path: solfatara.Path,
											Size: solfatara.Size,
											Md5sum: solfatara.Md5sum,
											IsDir: solfatara.IsDir,
										},
									},
									nuovo.Name: []*FileDetailsMap{
										&FileDetailsMap{
											Id: nuovo.Id,
											Name: nuovo.Name,
											Path: nuovo.Path,
											Size: nuovo.Size,
											Md5sum: nuovo.Md5sum,
											IsDir: nuovo.IsDir,
										},
									},
									averna.Name: []*FileDetailsMap{
										&FileDetailsMap{
											Id: averna.Id,
											Name: averna.Name,
											Path: averna.Path,
											Size: averna.Size,
											Md5sum: averna.Md5sum,
											IsDir: averna.IsDir,
										},
										&FileDetailsMap{
											Id: averna.Id,
											Name: averna.Name,
											Path: averna.Path,
											Size: averna.Size,
											Md5sum: averna.Md5sum,
											IsDir: averna.IsDir,
										},
									},
								},
							},
						},
					},
				},
			},
			sicilianVolcanoes.Name: []*FileDetailsMap{
				&FileDetailsMap{
					Id: sicilianVolcanoes.Id,
					Name: sicilianVolcanoes.Name,
					Path: sicilianVolcanoes.Path,
					Size: sicilianVolcanoes.Size,
					Md5sum: sicilianVolcanoes.Md5sum,
					IsDir: sicilianVolcanoes.IsDir,
					Children: FileTreeMap{
						etna.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: etna.Id,
								Name: etna.Name,
								Path: etna.Path,
								Size: etna.Size,
								Md5sum: etna.Md5sum,
								IsDir: etna.IsDir,
							},
						},
					},
				},
				&FileDetailsMap{
					Id: sicilianVolcanoes.Id,
					Name: sicilianVolcanoes.Name,
					Path: sicilianVolcanoes.Path,
					Size: sicilianVolcanoes.Size,
					Md5sum: sicilianVolcanoes.Md5sum,
					IsDir: sicilianVolcanoes.IsDir,
					Children: FileTreeMap{
						etna.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: etna.Id,
								Name: etna.Name,
								Path: etna.Path,
								Size: etna.Size,
								Md5sum: etna.Md5sum,
								IsDir: etna.IsDir,
							},
						},
					},
				},
			},
			volcanicIslands.Name: []*FileDetailsMap{
				&FileDetailsMap{
					Id: volcanicIslands.Id,
					Name: volcanicIslands.Name,
					Path: volcanicIslands.Path,
					Size: volcanicIslands.Size,
					Md5sum: volcanicIslands.Md5sum,
					IsDir: volcanicIslands.IsDir,
					Children: FileTreeMap{
						stromboli.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: stromboli.Id,
								Name: stromboli.Name,
								Path: stromboli.Path,
								Size: stromboli.Size,
								Md5sum: stromboli.Md5sum,
								IsDir: stromboli.IsDir,
							},
						},
						volcano.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: volcano.Id,
								Name: volcano.Name,
								Path: volcano.Path,
								Size: volcano.Size,
								Md5sum: volcano.Md5sum,
								IsDir: volcano.IsDir,
							},
						},
					},
				},
				&FileDetailsMap{
					Id: volcanicIslands.Id,
					Name: volcanicIslands.Name,
					Path: volcanicIslands.Path,
					Size: volcanicIslands.Size,
					Md5sum: volcanicIslands.Md5sum,
					IsDir: volcanicIslands.IsDir,
					Children: FileTreeMap{
						stromboli.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: stromboli.Id,
								Name: stromboli.Name,
								Path: stromboli.Path,
								Size: stromboli.Size,
								Md5sum: stromboli.Md5sum,
								IsDir: stromboli.IsDir,
							},
						},
						volcano.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: volcano.Id,
								Name: volcano.Name,
								Path: volcano.Path,
								Size: volcano.Size,
								Md5sum: volcano.Md5sum,
								IsDir: volcano.IsDir,
							},
						},
					},
				},
				&FileDetailsMap{
					Id: volcanicIslands.Id,
					Name: volcanicIslands.Name,
					Path: volcanicIslands.Path,
					Size: volcanicIslands.Size,
					Md5sum: volcanicIslands.Md5sum,
					IsDir: volcanicIslands.IsDir,
					Children: FileTreeMap{
						stromboli.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: stromboli.Id,
								Name: stromboli.Name,
								Path: stromboli.Path,
								Size: stromboli.Size,
								Md5sum: stromboli.Md5sum,
								IsDir: stromboli.IsDir,
							},
						},
						volcano.Name: []*FileDetailsMap{
							&FileDetailsMap{
								Id: volcano.Id,
								Name: volcano.Name,
								Path: volcano.Path,
								Size: volcano.Size,
								Md5sum: volcano.Md5sum,
								IsDir: volcano.IsDir,
							},
						},
					},
				},
			},
		},
		duplicates: []*FileDetailsMap{
			&FileDetailsMap{
				Id: averna.Id,
				Name: averna.Name,
				Path: "/" + italianMainlandVolcanoes.Name +
					"/" + campiflegrei.Name,
				Size: averna.Size,
				Md5sum: averna.Md5sum,
				IsDir: averna.IsDir,
			},
			&FileDetailsMap{
				Id: solfatara.Id,
				Name: solfatara.Name,
				Path: "/" + italianMainlandVolcanoes.Name +
					"/" + campiflegrei.Name,
				Size: solfatara.Size,
				Md5sum: solfatara.Md5sum,
				IsDir: solfatara.IsDir,
			},
			&FileDetailsMap{
				Id: solfatara.Id,
				Name: solfatara.Name,
				Path: "/" + italianMainlandVolcanoes.Name +
					"/" + campiflegrei.Name,
				Size: solfatara.Size,
				Md5sum: solfatara.Md5sum,
				IsDir: solfatara.IsDir,
			},
			&FileDetailsMap{
				Id: vesuvious.Id,
				Name: vesuvious.Name,
				Path: "/" + italianMainlandVolcanoes.Name,
				Size: vesuvious.Size,
				Md5sum: vesuvious.Md5sum,
				IsDir: vesuvious.IsDir,
			},
			&FileDetailsMap{
				Id: sicilianVolcanoes.Id,
				Name: sicilianVolcanoes.Name,
				Path: sicilianVolcanoes.Path,
				Size: sicilianVolcanoes.Size,
				Md5sum: sicilianVolcanoes.Md5sum,
				IsDir: sicilianVolcanoes.IsDir,
				Children: FileTreeMap{
					etna.Name: []*FileDetailsMap{
						&FileDetailsMap{
							Id: etna.Id,
							Name: etna.Name,
							Path: etna.Path,
							Size: etna.Size,
							Md5sum: etna.Md5sum,
							IsDir: etna.IsDir,
						},
					},
				},
			},
			&FileDetailsMap{
				Id: volcanicIslands.Id,
				Name: volcanicIslands.Name,
				Path: volcanicIslands.Path,
				Size: volcanicIslands.Size,
				Md5sum: volcanicIslands.Md5sum,
				IsDir: volcanicIslands.IsDir,
				Children: FileTreeMap{
					stromboli.Name: []*FileDetailsMap{
						&FileDetailsMap{
							Id: stromboli.Id,
							Name: stromboli.Name,
							Path: stromboli.Path,
							Size: stromboli.Size,
							Md5sum: stromboli.Md5sum,
							IsDir: stromboli.IsDir,
						},
					},
					volcano.Name: []*FileDetailsMap{
						&FileDetailsMap{
							Id: volcano.Id,
							Name: volcano.Name,
							Path: volcano.Path,
							Size: volcano.Size,
							Md5sum: volcano.Md5sum,
							IsDir: volcano.IsDir,
						},
					},
				},
			},
			&FileDetailsMap{
				Id: volcanicIslands.Id,
				Name: volcanicIslands.Name,
				Path: volcanicIslands.Path,
				Size: volcanicIslands.Size,
				Md5sum: volcanicIslands.Md5sum,
				IsDir: volcanicIslands.IsDir,
				Children: FileTreeMap{
					stromboli.Name: []*FileDetailsMap{
						&FileDetailsMap{
							Id: stromboli.Id,
							Name: stromboli.Name,
							Path: stromboli.Path,
							Size: stromboli.Size,
							Md5sum: stromboli.Md5sum,
							IsDir: stromboli.IsDir,
						},
					},
					volcano.Name: []*FileDetailsMap{
						&FileDetailsMap{
							Id: volcano.Id,
							Name: volcano.Name,
							Path: volcano.Path,
							Size: volcano.Size,
							Md5sum: volcano.Md5sum,
							IsDir: volcano.IsDir,
						},
					},
				},
			},
		},
	},
}

func TestGroupByName(t *testing.T) {
	for _, pair := range tests {
		got := GroupByName(pair.input)
		if !compareFileTreeMaps(got, pair.grouped) {
			t.Error("Wrong Group By Name")
		}
	}
}

func compareFileTreeMaps(got, expected FileTreeMap) bool {
	if got == nil && expected == nil {
		return true
	}
	if len(got) != len(expected) {
		return false
	}
	for name, fileDetailsMapList := range expected {
		if len(got[name]) != len(fileDetailsMapList) {
			return false
		}
		for i := range fileDetailsMapList {
			if !compareFileDetailsMap(got[name][i], fileDetailsMapList[i]) {
				return false
			}
		}
	}
	return true
}

func compareFileDetailsMap(gotFileDetailsMap,
	expectedFileDetailsMap *FileDetailsMap) bool {

	if gotFileDetailsMap == nil && expectedFileDetailsMap == nil {
		return true
	}
	if expectedFileDetailsMap == nil {
		color.Yellow("Unexpected file \"%+v\"", gotFileDetailsMap)
		return false
	} else if gotFileDetailsMap == nil {
		color.Yellow("Missing file \"%+v\"", expectedFileDetailsMap)
		return false
	}
	// Compare IDs only if they are defined
	if (gotFileDetailsMap.Id != "" && expectedFileDetailsMap.Id != "") &&
	(gotFileDetailsMap.Id != expectedFileDetailsMap.Id) {
		color.Yellow("Expected ID is different. Got \"%+v\", Expected \"%+v\"",
			gotFileDetailsMap, expectedFileDetailsMap)
		return false
	}
	// Compare names
	if gotFileDetailsMap.Name != expectedFileDetailsMap.Name {
		color.Yellow("Expected name is different. Got \"%s\", Expected \"%s\"",
			gotFileDetailsMap.Name, expectedFileDetailsMap.Name)
		return false
	}
	// Compare paths
	if gotFileDetailsMap.Path != expectedFileDetailsMap.Path {
		color.Yellow("Expected path is different. Got \"%s\", Expected \"%s\""+
			"\n%+v\n%+v\n",
			gotFileDetailsMap.Path, expectedFileDetailsMap.Path,
			gotFileDetailsMap, expectedFileDetailsMap,
		)
		return false
	}
	// Compare sizes
	if gotFileDetailsMap.Size != expectedFileDetailsMap.Size {
		color.Yellow("Expected size is different. Got \"%s\", Expected \"%s\"",
			gotFileDetailsMap.Size, expectedFileDetailsMap.Size)
		return false
	}
	// Compare md5sums of two files
	if gotFileDetailsMap.Md5sum != expectedFileDetailsMap.Md5sum {
		color.Yellow("Expected md5 is different. Got \"%s\", Expected\"%s\"",
			gotFileDetailsMap.Md5sum, expectedFileDetailsMap.Md5sum)
		return false
	}
	// Are both files dirs?
	if gotFileDetailsMap.IsDir != expectedFileDetailsMap.IsDir {
		color.Yellow("Expected file is not a dir. Got \"%s\", Expected \"%s\"",
			gotFileDetailsMap.IsDir, expectedFileDetailsMap.IsDir)
		return false
	}

	// If everything above is equal, let's compare children items.
	return compareFileTreeMaps(
		gotFileDetailsMap.Children,
		expectedFileDetailsMap.Children,
	)
}

func TestFindDuplicates(t *testing.T) {
	for _, pair := range tests {
		groupedFileTreeMap := GroupByName(pair.input)
		got := FindDuplicates("/", groupedFileTreeMap)
		if !compareFileTreeMaps(
			FileTreeMap{ "duplicates": got },
			FileTreeMap{ "duplicates": pair.duplicates },
		) {
			t.Error("Wrong Find Duplicates")
		}
	}
}

func TestMain(m *testing.M) {
	IsTesting = true
	flag.Parse()
	InitLoggers("debug", "warn", "error")
	os.Exit(m.Run())
}
