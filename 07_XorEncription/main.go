package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type config struct {
	inputPath  string
	outputPath string
	key        []byte
	op         string
}

func main() {
	config, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputFile, outputFile, err := openFiles(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = inputFile.Close()
		// output file opens for writing, that why we check error when close, unlike to input file, which open for reading
		if err := outputFile.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	if err := do(inputFile, outputFile, config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseArgs() (*config, error) {
	inputPath := flag.String("f", "", "file with original message to encrypt or encrypted message to decrypt")
	key := flag.String("k", "", "key for encryption/decryption")
	outputPath := flag.String("o", "", "file for result encryption/decryption")
	op := flag.String("t", "", "type of operation [encrypt|decrypt]")
	flag.Parse()
	if *inputPath == "" {
		return nil, errors.New("input file not specified")
	}
	if *outputPath == "" {
		return nil, errors.New("output file not specified")
	}
	if *key == "" {
		return nil, errors.New("key is empty")
	}
	if *op != "encrypt" && *op != "decrypt" {
		return nil, fmt.Errorf("unknown opeartion %q, use encrypt or decrypt", *op)
	}
	return &config{
		inputPath:  *inputPath,
		outputPath: *outputPath,
		key:        []byte(*key),
		op:         *op,
	}, nil
}

func openFiles(config *config) (inputFile *os.File, outputFile *os.File, err error) {
	inputFile, err = os.Open(config.inputPath)
	if err != nil {
		return nil, nil, err
	}
	outputFile, err = os.Create(config.outputPath)
	if err != nil {
		inputFile.Close()
		return nil, nil, err
	}
	return
}

func do(inputFile *os.File, outputFile *os.File, config *config) error {
	switch config.op {
	case "encrypt":
		return Encrypt(inputFile, outputFile, config.key)
	case "decrypt":
		return Decrypt(inputFile, outputFile, config.key)
	default:
		return errors.New("unknown operation type, use encrypt or decrypt")
	}
}
