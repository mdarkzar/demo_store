run:
	cd ./scripts/ && ./start_docker.sh
test:
	cd ./scripts/ && ./run_test_all.sh
clean:
	cd ./docker && rm -rf volumes/