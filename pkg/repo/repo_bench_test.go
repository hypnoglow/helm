package repo

import "testing"

func BenchmarkRepoFile_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repoFile().Remove("first")
	}
}

func repoFile() *RepoFile {
	repo := NewRepoFile()
	repo.Add(
		&Entry{
			Name:  "first",
			URL:   "https://example.com/first",
			Cache: "first-index.yaml",
		},
		&Entry{
			Name:  "second",
			URL:   "https://example.com/second",
			Cache: "second-index.yaml",
		},
		&Entry{
			Name:  "third",
			URL:   "https://example.com/third",
			Cache: "third-index.yaml",
		},
		&Entry{
			Name:  "fourth",
			URL:   "https://example.com/fourth",
			Cache: "fourth-index.yaml",
		},
		&Entry{
			Name:  "fifth",
			URL:   "https://example.com/fifth",
			Cache: "fifth-index.yaml",
		},
		&Entry{
			Name:  "sixth",
			URL:   "https://example.com/sixth",
			Cache: "sixth-index.yaml",
		},
		&Entry{
			Name:  "seventh",
			URL:   "https://example.com/seventh",
			Cache: "seventh-index.yaml",
		},
		&Entry{
			Name:  "eighth",
			URL:   "https://example.com/eighth",
			Cache: "eighth-index.yaml",
		},
	)
	return repo
}
