package main

import (
	"app/ipfs-json-upload-poc/nft"
	"bytes"
	"encoding/json"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io"
	"os"
	"strconv"
)

func main() {

	fmt.Println("IPFS LOAD......")
	dirNm := "nfts-poc"
	createJSON(dirNm)
	sh := shell.NewShell("localhost:5001")
	code, err := sh.AddDir(dirNm)
	if err != nil {
		check(err)
	}
	fmt.Println(code)
	//sh.FilesRm()
	//hash : QmcHcwh5U8A3G4EuycidAsA1Nyo3AqMuUjrQQDT1TyJcCE
}

func createJSON(dirName string) {

	e := os.Mkdir(dirName, os.FileMode.Perm(0775))
	check(e)
	listJson := getMockMetadata()

	for i, metadataNFT := range listJson {

		nmFile := strconv.Itoa(i)
		f, errCreate := os.Create(dirName + "/" + nmFile + ".json")
		check(errCreate)
		defer f.Close()
		var buffer bytes.Buffer
		errEncode := prettyEncode(metadataNFT, &buffer)
		check(errEncode)
		_, errWriter := f.Write(buffer.Bytes())
		check(errWriter)
	}
}

func prettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}
func getMockMetadata() []nft.MetadataNFT {

	var mocks []nft.MetadataNFT
	i := 1
	for i < 100 {
		stri := strconv.Itoa(i)
		var mock nft.MetadataNFT
		mock.Title = "Asset Metadata " + stri
		mock.Type = "NFT"
		mock.Properties.Name.Type = "string"
		mock.Properties.Name.Description = "Identifies the asset to which this NFT represents " + stri
		mock.Properties.Description.Description = "Describes the asset to which this NFT represents " + stri
		mock.Properties.Description.Type = "string"
		mock.Properties.Image.Type = "string"
		mock.Properties.Image.Description = "https://i.pinimg.com/564x/e3/0b/ee/e30beeed451d4ef26ccf9aa007783580.jpg"

		mock.Attributes = append(mock.Attributes, nft.Attribute{
			TraitType: "Background",
			Value:     "Blue",
		})

		mock.Attributes = append(mock.Attributes, nft.Attribute{
			TraitType: "Head",
			Value:     "Hat",
		})

		mock.Attributes = append(mock.Attributes, nft.Attribute{
			TraitType:   "Level",
			Value:       stri,
			DisplayType: "number",
		})

		mock.Attributes = append(mock.Attributes, nft.Attribute{
			TraitType:   "Health",
			Value:       stri,
			DisplayType: "number",
		})

		i++
		mocks = append(mocks, mock)
	}
	return mocks
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
