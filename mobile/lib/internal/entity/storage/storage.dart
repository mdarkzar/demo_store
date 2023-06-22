import 'package:freezed_annotation/freezed_annotation.dart';

part 'storage.freezed.dart';
part 'storage.g.dart';

@freezed
abstract class Storage with _$Storage {
  const factory Storage(
    int id,
    String name,
  ) = _Storage;
  factory Storage.fromJson(Map<String, dynamic> json) =>
      _$StorageFromJson(json);
}

@freezed
abstract class StorageList with _$StorageList {
  const factory StorageList(List<Storage>? storageList) = _StorageList;
  factory StorageList.fromJson(Map<String, dynamic> json) =>
      _$StorageListFromJson(json);
}
