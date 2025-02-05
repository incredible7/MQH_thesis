# MQH_thesis

## Structure

MQH_THESIS
|- data
|  |- datasets
|  |- results
|  |- ...
|- scripts
|  |- download_datasets.py
|  |- ...
|- exhaustive.go
|- exhaustive_test.go
|- README.md
|- ...

## Benchmarking datasets
| Data Sets | # Data Points | # Dim | # Queries | Data Size | Data Type |
| --------- | ------------- | ----- | --------- | --------- | --------- |
| Music     | 1,000,000     | 100   | 100       | 386 MB    | Rating    |
| GloVe     | 1,183,514     | 100   | 100       | 460 MB    | Text      |
| Sift      | 985,462       | 128   | 100       | 485 MB    | Image     |
| UKBench   | 1,097,907     | 128   | 100       | 541 MB    | Image     |
| Tiny      | 1,000,000     | 384   | 100       | 1.5 GB    | Image     |
| Msong     | 992,272       | 420   | 100       | 1.6 GB    | Audio     |
| NUSW      | 268,643       | 500   | 100       | 514 MB    | Image     |
| Cifar-10  | 50,000        | 512   | 100       | 98 MB     | Image     |
| Sun       | 79,106        | 512   | 100       | 155 MB    | Image     |
| LabelMe   | 181,093       | 512   | 100       | 355 MB    | Image     |
| Gist      | 982,694       | 960   | 100       | 3.6 GB    | Image     |
| Enron     | 94,987        | 1,369 | 100       | 497 MB    | Text      |
| Trevi     | 100,900       | 4,096 | 100       | 1.6 GB    | Image     |
| P53       | 31,153        | 5,408 | 100       | 643 MB    | Biology   |
| Deep100M  | 100,000,000   | 96    | 100       | 36.1 GB   | Image     |
| Sift100M  | 99,986,452    | 128   | 100       | 48.0 GB   | Image     |


## To fetch the datasets from the BC-Tree-Datasets folder provided by Huang Qiang
Be aware that the datasets are large, so the download may take a while.
```bash
pip install gdown
python scripts/download_datasets.py
```

## To run main function in GT mode?

```bash
go run cmd/benchmark/main.go Cifar 50000 512 100 100
```

Exhaustive searchPQ is inspired from (this link)[https://pkg.go.dev/container/heap#example-package-PriorityQueue].

