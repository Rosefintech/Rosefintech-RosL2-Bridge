/*
@Author :
@group :Rosefintech
@Description :
@Data: 2021/4/28 10:34
*/

package tools

import (
	"strings"
)

func ZokCompile(name string) error {
	_, err := ExecLinuxApp("zokrates compile -i " + name)
	return err
}


func ComputeWitness(abiFile, arguments, inputFile, outPutFile string) (proof []byte, err error) {
	bytes, err := ExecLinuxApp(strings.Join([]string{"./zokrates compute-witness -s", abiFile, "-a", arguments, "-i", inputFile, "-o", outPutFile}, " "))
	if err != nil {
		return
	}
	Logger.Debug(string(bytes))
	return
}


//generate a zero-knowledge proof
func GenerateProof(inputFile, provingScheme string) (proof []byte, err error) {
	_, err = ExecLinuxApp(strings.Join([]string{"./zokrates generate-proof -i", inputFile, "-s", provingScheme}, " "))
	proof, err = ReadFile("proof.json")
	return
}

