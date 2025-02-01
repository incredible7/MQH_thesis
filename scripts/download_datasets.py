import os
import gdown

# create datasets directory if it doesn't exist
os.makedirs("../data/datasets", exist_ok=True)

# list of all files within the BC-Tree-Datasets folder and their IDs
# currently only Cifar-10 is downloaded
datasets = [
    ("Cifar.ds","1uUEOCqvxhvM8qkRkucy3_Iln02lCOcxW"),
    ("Cifar.gt","1g9jKP6PFyBOpOupBLXMqT_Fq8lC3RO43"),
    ("Cifar.q","1h4d2Uid_e3cUb6ywkwOy2HBIIyL-Hpui"),
    # ("Enron.ds","1BB8Sdkf1d0Oi3bX7togUDsOJJKtqVDPp"),
    # ("Enron.gt","1aEyHG8oEHWgyAojSEjELLeUnPuajEfVd"),
    # ("Enron.q","1e0BkBQjtK8AkAKRMqpsH1GRZjZpgXtTC"),
    # ("Gist.ds","12Nproek3QPQB_EgxnFhWaWLIZKTCx1lf"),
    # ("Gist.gt","1ylyhXTn6fh_bhmwMJS4-sPtUvoS9Xuq1"),
    # ("Gist.q","1i10kqIlQJ8DkiStxiMOd5Q8bBPp4bhJG"),
    # ("GloVe100.ds","1C-OXbGJjfouYhqvhPbL57L5-uRAJoiGK"),
    # ("GloVe100.gt","1iEn4hUkPBpb3xDoFHas0l4RN87swEdr-"),
    # ("GloVe100.q","1aH9fWAcrd6oq2l8VbhbltZtD82Gfizmv"),
    # ("LabelMe.ds","1PRksjQoBbtUNkWusq39hU2_7L7qdjcw1"),
    # ("LabelMe.gt","1fJRHSdadziHELuScUE-ak7riCNATMkri"),
    # ("LabelMe.q","1n0LwjWpUumTensda0uotqyvc8JB4jU4X"),
    # ("Msong.ds","1aZuXnq-TWyed3GoXeVP4uNEI9EMNCX4G"),
    # ("Msong.gt","1k4SauiWzD9k1_Wh3GKVAw2hAVf5jLEbg"),
    # ("Msong.q","15NC1bWTpG_LdCM_QJPocj33WJb2lI8UQ"),
    # ("Music.ds","1dmK9Iw0Gq-259hknL3oFUw6C2gclNBje"),
    # ("Music.gt","1hhs_WRs6ZlkL29PH0ep0Xnor_uWA_tsB"),
    # ("Music.q","1P4TMfSH0iWMo6NPRDcqytsUV1A_Dtk9k"),
    # ("NUSW.ds","1VDtVCcobLLPbuRkjlK7JM1SPWh2vNzNW"),
    # ("NUSW.gt","1LEEKnxFSpFF84eI_DfEV6RF_udt9fZ93"),
    # ("NUSW.q","1IWySoIOQWRNHn0eS6GuPAmQvLA2NsV3x"),
    # ("P53.ds","1wmXSnRbUCOgqVs3J9e1glD7mGs5LijFC"),
    # ("P53.gt","1m_CSjmNpZvyo-A9Cl8Ykab6n-VvKFZqK"),
    # ("P53.q","1IsvKVLiillOIYgJdvjrVLhhaD1bZEh4S"),
    # ("Sift.ds","1PRSMYBRqkt1gJcVBP5wS8jGg0ScjC01q"),
    # ("Sift.gt","1j7Dz7uAz8Vsr3tAMLJSqvcnvNLBRl-A4"),
    # ("Sift.q","1RZwwCVTH2cT3xCTMVagE0jXytHUZYAFC"),
    # ("Sun.ds","1fZ3TQecEgt3tKI9b4AhtJ6kM3uPwPTYD"),
    # ("Sun.gt","1KnJ1Z7E9Zm59RakOQzw5bq8gKnFx99N-"),
    # ("Sun.q","10BmU77N4fFyqqUZ8LKJ2SUuwNVPblR_h"),
    # ("Tiny1M.ds","1dNOQlZi7jb1S5os-b-fhcwAVEOQw0OXL"),
    # ("Tiny1M.gt","1nOKmDYR3hupWaA51onBpO80SFcpN1SqH"),
    # ("Tiny1M.q","1Q7UDBwpfkrmnZPO46dyzcaCgjlKRyYKk"),
    # ("Trevi.ds","1FaGJH8Gvv8agglECLIM5zDyoc7jq0-XN"),
    # ("Trevi.gt","15ZpB0jZXtjmNwVe19tWHqhgEvT6SxEFb"),
    # ("Trevi.q","1TAxFXLmBfCl5C4yHzgcCmvTfg-EQXfSk"),
    # ("UKBench.ds","17Ve5rRxfNoShmg-O9HRbRkXVQCj3i-Sz"),
    # ("UKBench.gt","1e-mXDGWXD0xrmOCNClH7Tumd03_xFx2t"),
    # ("UKBench.q","1Ecg8dB31HsHHnHyUG0oMPAkpJ31n6hWa")
]

def download_file(file_name, file_id):
    output_dir = "data/datasets"
    output_path = os.path.join(output_dir, file_name)
    if not os.path.exists(output_path):
        url = f"https://drive.google.com/uc?id={file_id}"
        try:
            gdown.download(url, output_path, quiet=False, fuzzy=True)
            print(f"Successfully downloaded {file_name}")
        except Exception as e:
            print(f"Error downloading {file_name}: {str(e)}")
    else:
        print(f"Skipping {file_name} (already exists)")

def main():
    print("Starting downloads...")
    for file_name, file_id in datasets:
        download_file(file_name, file_id)
    print("Download complete!")

if __name__ == "__main__":
    main()