set -x

version=${1:?"Version must be provided"} 
echo $version

git fetch git@github.com:golang/crypto.git $version
git checkout FETCH_HEAD
git cherry-pick master..pbkdf2_switch  # pbkdf changes
git cherry-pick master..hkdf_switch  # hkdf changes
git tag -a $version-kastenhq -m "$version kastenhq release"
git push origin $version-kastenhq
