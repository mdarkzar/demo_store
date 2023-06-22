// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'storage.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_Storage _$$_StorageFromJson(Map<String, dynamic> json) => _$_Storage(
      json['id'] as int,
      json['name'] as String,
    );

Map<String, dynamic> _$$_StorageToJson(_$_Storage instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
    };

_$_StorageList _$$_StorageListFromJson(Map<String, dynamic> json) =>
    _$_StorageList(
      (json['storageList'] as List<dynamic>?)
          ?.map((e) => Storage.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$$_StorageListToJson(_$_StorageList instance) =>
    <String, dynamic>{
      'storageList': instance.storageList,
    };
