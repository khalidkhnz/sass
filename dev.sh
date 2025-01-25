echo RUN GO_BLOG
cd ./go-blog
make run &

echo RUN GO_ECOM 
cd ../go-ecom
make run &

echo RUN GO_GATEWAY
cd ../go-gateway
make run &

echo RUN GO_SASS
cd ../go-sass
make run &

cd ..

wait
