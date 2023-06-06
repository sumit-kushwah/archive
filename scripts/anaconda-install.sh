# download anaconda script file
# wget https://repo.anaconda.com/archive/Anaconda3-2023.03-1-Linux-x86_64.sh

# hash value check here https://docs.anaconda.com/free/anaconda/reference/hashes/Anaconda3-2023.03-1-Linux-x86_64.sh-hash/

expected_hash="95102d7c732411f1458a20bdf47e4c1b0b6c8a21a2edfe4052ca370aaae57bab"

# Calculate the hash value of the file
actual_hash=$(shasum -a 256 Anaconda3-2023.03-1-Linux-x86_64.sh | awk '{print $1}')

if [[ "$actual_hash" == "$expected_hash" ]]; then
    # Run the script
    bash Anaconda3-2023.03-1-Linux-x86_64.sh
else
    echo "Hash mismatch. The file does not match the expected hash value."
fi