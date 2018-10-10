package models

import "gopkg.in/mgo.v2/bson"

//Login - Data type for user login
type Login struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

//GuruPrimary - Data type for teacher primary data
type GuruPrimary struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	NIK        string        `bson:"nik" json:"nik"`
	Nama       string        `bson:"nama" json:"nama"`
	Jenkel     string        `bson:"jenkel" json:"jenkel"`
	Email      string        `bson:"email" json:"email"`
	Kota       string        `bson:"kota" json:"kota"`
	TglLahir   string        `bson:"tgl_lahir" json:"tgl_lahir"`
	Telepon    string        `bson:"telepon" json:"telepon"`
	Password   string        `bson:"password" json:"password"`
	Pendidikan pendidikan    `bson:"pendidikan" json:"pendidikan"`
	Info       info          `bson:"info" json:"info"`
}

//SiswaPrimary - Data type for student primary data
type SiswaPrimary struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Nama     string        `bson:"nama" json:"nama"`
	Jenkel   string        `bson:"jenkel" json:"jenkel"`
	Email    string        `bson:"email" json:"email"`
	Kota     string        `bson:"kota" json:"kota"`
	TglLahir string        `bson:"tgl_lahir" json:"tgl_lahir"`
	Telepon  string        `bson:"telepon" json:"telepon"`
	Password string        `bson:"password" json:"password"`
}

type sekolah struct {
	Nama    string `bson:"nama" json:"nama"`
	Jurusan string `bson:"jurusan" json:"jurusan"`
	Masuk   string `bson:"tahun_masuk" json:"tahun_masuk"`
	Lulus   string `bson:"tahun_lulus" json:"tahun_lulus"`
}

type pendidikan struct {
	ID  bson.ObjectId `bson:"_id" json:"id"`
	SD  sekolah       `bson:"sd" json:"sd"`
	SMP sekolah       `bson:"smp" json:"smp"`
	SMA sekolah       `bson:"sma" json:"sma"`
	PT  sekolah       `bson:"pt" json:"pt"`
}

type info struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Deskripsi string        `bson:"deskripsi" json:"deskripsi"`
	Prestasi  string        `bson:"prestasi" json:"prestasi"`
	Bidang    []string      `bson:"bidang" json:"bidang"`
}
