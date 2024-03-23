set -x

version=${1:?"Version must be provided"} 
echo $version

git fetch git@github.com:golang/crypto.git $version
git checkout FETCH_HEAD
git cherry-pick 12b6180  # pbkdf changes
git cherry-pick 5d64f3c  # hkdf changes
git tag -a $version-kastenhq -m "$version kastenhq release"
git push origin $version-kastenhq
