package invasion

import "testing"

func TestReadCityMapFile_ValidFilePath(t *testing.T) {
	path := "./testdata/map.txt"

	cityMap, err := ReadWorldMapFile(path)
	if err != nil {
		t.Errorf("An error occurred while reading the file: %v", err)
	}

	if cityMap["Qu-ux"].connections["north"] != "Baz" {
		t.Errorf("Error in reading the citymap")
	}
}

func TestReadCityMapFile_InValidFilePath(t *testing.T) {
	invalidPath := "./testdata/notafile.txt"

	_, err := ReadWorldMapFile(invalidPath)
	if err == nil {
		t.Errorf("Expected error for invalid file path")
	}
}

func TestReadCityMapFile_FileWithEmptyLines(t *testing.T) {
	wrongFormatPath := "./wrongFormatCityMap.txt"

	_, err := ReadWorldMapFile(wrongFormatPath)
	if err == nil {
		t.Errorf("Expected error for file with wrong format")
	}
}
