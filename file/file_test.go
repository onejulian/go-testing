package file

import ("testing"

"github.com/stretchr/testify/assert")

func TestCreateFile(t *testing.T){
	assert := assert.New(t)
	file := File{}
	
	err := file.CreateFile("test.txt", "../test", "")
	assert.NotNil(err)
	
	err = file.CreateFile("test.txt", "", "test")
	assert.NotNil(err)
	
	err = file.CreateFile("", "../test", "test")
	assert.NotNil(err)
	err = file.CreateFile("test.txt", "../test", "test")
	assert.Nil(err)
}

func TestReadFile(t *testing.T){
	assert := assert.New(t)
	file := File{}
	content, err := file.ReadFile("test.txt", "../test")
	assert.Nil(err)
	assert.Equal("test", content)

	_, err = file.ReadFile("", "../test")
	assert.NotNil(err)

	_, err = file.ReadFile("test.txt", "")
	assert.NotNil(err)
}

func TestUpdateFile(t *testing.T){
	assert := assert.New(t)
	file := File{}
	err := file.UpdateFile("test.txt", "../test", "test2")
	assert.Nil(err)

	err = file.UpdateFile("", "../test", "test2")
	assert.NotNil(err)

	err = file.UpdateFile("test.txt", "../test", "")
	assert.NotNil(err)
}

func TestDeleteFile(t *testing.T){
	assert := assert.New(t)
	file := File{}
	
	err := file.DeleteFile("", "../test")
	assert.NotNil(err)
	
	err = file.DeleteFile("test.txt", "")
	assert.NotNil(err)
	
	err = file.DeleteFile("test.txt", "../test")
	assert.Nil(err)
}