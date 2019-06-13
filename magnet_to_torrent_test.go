package magnettotorrent

import "testing"

func TestGetTorrent(t *testing.T){

	hash := "B415C913643E5FF49FE37D304BBB5E6E11AD5101"
	expected := "http://itorrents.org/torrent/B415C913643E5FF49FE37D304BBB5E6E11AD5101.torrent"
	result := GetTorrent(hash)

	t.Log(isValidURL(expected))
	// if result != expected {
	// 	t.Fatalf("\nExpected :	%s \nResult :	%s",expected,result)
	// }

	wrongHash := "PIKACHU"
	expected = ""

	result = GetTorrent(wrongHash)
	if result != expected {
		t.Fatalf("\nExpected :		 %s \nResult :		 %s",expected,result)
	}

	


}