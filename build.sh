echo GO_BLOG
cd ./go-blog
make build

echo GO_ECOM
cd ../go-ecom
make build

echo GO_GATEWAY
cd ../go-gateway
make build

echo GO_SASS
cd ../go-sass
make build

cd ..
ls

git add .
git commit -m "DEPLOY:BUILD"