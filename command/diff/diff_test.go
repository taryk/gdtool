package diff

import (
	"testing"
	"github.com/fatih/color"
	"os"
	"flag"

	. "github.com/taryk/gdtool/core"
)

type FileLists struct {
	remote, local FileList
}

type testpair struct {
	title string
	input, expected FileLists
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
		title: "No unique stuff when comparing emtpy lists",
		input: FileLists{
			remote: FileList{},
			local: FileList{},
		},
		expected: FileLists{
			remote: nil,
			local: nil,
		},
	},
	{
		title: "No unique stuff when comparing emtpy directories",
		input: FileLists{
			remote: FileList{empty},
			local: FileList{empty},
		},
		expected: FileLists{
			remote: nil,
			local: nil,
		},
	},
	{
		title: "Remote empty directory is unique",
		input: FileLists{
			remote: FileList{empty},
			local: FileList{},
		},
		expected: FileLists{
			remote: FileList{empty},
			local: nil,
		},
	},
	{
		title: "Lists of files are the same",
		input: FileLists{
			remote: FileList{vesuvious, solfatara, etna, volcano},
			local: FileList{vesuvious, solfatara, etna, volcano},
		},
		expected: FileLists{
			remote: nil,
			local: nil,
		},
	},
	{
		title: "Different remote and local files",
		input: FileLists{
			remote: FileList{etna, solfatara},
			local: FileList{vesuvious, volcano},
		},
		expected: FileLists{
			remote: FileList{etna, solfatara},
			local: FileList{vesuvious, volcano},
		},
	},
	{
		title: "All remote files are unique",
		input: FileLists{
			remote: FileList{etna, solfatara, vesuvious, volcano},
			local: FileList{},
		},
		expected: FileLists{
			remote: FileList{etna, solfatara, vesuvious, volcano},
			local: nil,
		},
	},
	{
		title: "All local files are unique",
		input: FileLists{
			remote: FileList{},
			local: FileList{solfatara, vesuvious, volcano, etna},
		},
		expected: FileLists{
			remote: nil,
			local: FileList{solfatara, vesuvious, volcano, etna},
		},
	},
	{
		title: "Remote and local files have the same name but different md5",
		input: FileLists{
			remote: FileList{volcanicIslandsDir(stromboli)},
			local: FileList{volcanicIslandsDir(stromboliDifferentMd5sum)},
		},
		expected: FileLists{
			remote: FileList{volcanicIslandsDir(stromboli)},
			local: FileList{volcanicIslandsDir(stromboliDifferentMd5sum)},
		},
	},
	{
		title: "Remote and local nested file structures are the same",
		input: FileLists{
			remote: FileList{italianVolcanoes},
			local: FileList{italianVolcanoes},
		},
		expected: FileLists{
			remote: nil,
			local: nil,
		},
	},
	{
		title: "There are unique local nested files",
		input: FileLists{
			remote: FileList{partialItalianVolcanoes},
			local: FileList{italianVolcanoes},
		},
		expected: FileLists{
			remote: nil,
			local: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(solfatara, nuovo, averna),
					sicilianVolcanoesDir(etna),
					volcanicIslandsDir(stromboli),
				),
			},
		},
	},
	{
		title: "There are unique remote and local nested files",
		input: FileLists{
			remote: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(campiflegrei),
					partialVolcanicIslands,
				),
			},
			local: FileList{partialItalianVolcanoes},
		},
		expected: FileLists{
			remote: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(campiflegrei),
				),
			},
			local: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(vesuvious),
					sicilianVolcanoesDir(),
				),
			},
		},
	},
	{
		title: "Nested file duplicates get ignored",
		input: FileLists{
			remote: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(campiflegrei),
					italianMainlandVolcanoesDir(campiflegrei),
					partialVolcanicIslands,
					partialVolcanicIslands,
					partialVolcanicIslands,
				),
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(campiflegrei),
					italianMainlandVolcanoesDir(campiflegrei),
					italianMainlandVolcanoesDir(campiflegrei),
				),
			},
			local: FileList{
				partialItalianVolcanoes,
				partialItalianVolcanoes,
			},
		},
		expected: FileLists{
			remote: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(campiflegrei),
				),
			},
			local: FileList{
				italianVolcanoesDir(
					italianMainlandVolcanoesDir(vesuvious),
					sicilianVolcanoesDir(),
				),
			},
		},
	},
}

func TestCompare(t *testing.T) {
	for _, pair := range tests {
		uniqueRemote, uniqueLocal := Compare(
			pair.input.remote, pair.input.local,
		)
		if !compareFileLists(uniqueRemote, pair.expected.remote) {
			t.Errorf("%s\nremote:\n%s\n%s\n",
				pair.title,
				FileTreeStr("got:", uniqueRemote),
				FileTreeStr("expected:", pair.expected.remote))
		}
		if !compareFileLists(uniqueLocal, pair.expected.local) {
			t.Errorf("%s\nlocal:\n%s\n%s\n",
				pair.title,
				FileTreeStr("got:", uniqueLocal),
				FileTreeStr("expected:", pair.expected.local))
		}
	}
}

func compareFileLists(got, expected FileList) bool {
	if got == nil && expected == nil {
		return true
	}
	if len(got) != len(expected) {
		return false
	}
	for i := range expected {
		if !compareFileDetails(got[i], expected[i]) {
			return false
		}
	}
	return true
}

func compareFileDetails(gotFileDetails,
	expectedFileDetails *FileDetails) bool {

	if gotFileDetails == nil && expectedFileDetails == nil {
		return true
	}
	if expectedFileDetails == nil {
		color.Yellow("Unexpected file \"%+v\"", gotFileDetails)
		return false
	} else if gotFileDetails == nil {
		color.Yellow("Missing file \"%+v\"", expectedFileDetails)
		return false
	}
	// Compare IDs only if they are defined
	if (gotFileDetails.Id != "" && expectedFileDetails.Id != "") &&
	(gotFileDetails.Id != expectedFileDetails.Id) {
		color.Yellow("Expected ID is different. Got \"%+v\", Expected \"%+v\"",
			gotFileDetails, expectedFileDetails)
		return false
	}
	// Compare names
	if gotFileDetails.Name != expectedFileDetails.Name {
		color.Yellow("Expected name is different. Got \"%s\", Expected \"%s\"",
			gotFileDetails.Name, expectedFileDetails.Name)
		return false
	}
	// Compare paths
	if gotFileDetails.Path != expectedFileDetails.Path {
		color.Yellow("Expected path is different. Got \"%s\", Expected \"%s\"",
			gotFileDetails.Path, expectedFileDetails.Path)
		return false
	}
	// Compare sizes
	if gotFileDetails.Size != expectedFileDetails.Size {
		color.Yellow("Expected size is different. Got \"%s\", Expected \"%s\"",
			gotFileDetails.Size, expectedFileDetails.Size)
		return false
	}
	// Compare md5sums of two files
	if gotFileDetails.Md5sum != expectedFileDetails.Md5sum {
		color.Yellow("Expected md5 is different. Got \"%s\", Expected\"%s\"",
			gotFileDetails.Md5sum, expectedFileDetails.Md5sum)
		return false
	}
	// Are both files dirs?
	if gotFileDetails.IsDir != expectedFileDetails.IsDir {
		color.Yellow("Expected file is not a dir. Got \"%s\", Expected \"%s\"",
			gotFileDetails.IsDir, expectedFileDetails.IsDir)
		return false
	}

	// If everything above is equal, let's compare children items.
	return compareFileLists(
		gotFileDetails.Children,
		expectedFileDetails.Children,
	)
}

func TestMain(m *testing.M) {
	IsTesting = true
	flag.Parse()
	InitLoggers("debug", "warn", "error")
	os.Exit(m.Run())
}
