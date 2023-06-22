run:
	cd ./scripts/ && ./start_docker.sh
test:
	cd ./scripts/ && ./run_test_all.sh
clean:
	cd ./docker && rm -rf volumes/
build android:
	cd ./mobile && flutter pub get && flutter build apk --no-tree-shake-icons && cp build/app/outputs/flutter-apk/app-release.apk demo_store_v1.apk