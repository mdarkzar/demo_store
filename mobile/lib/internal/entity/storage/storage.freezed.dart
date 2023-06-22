// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'storage.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

Storage _$StorageFromJson(Map<String, dynamic> json) {
  return _Storage.fromJson(json);
}

/// @nodoc
mixin _$Storage {
  int get id => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $StorageCopyWith<Storage> get copyWith => throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $StorageCopyWith<$Res> {
  factory $StorageCopyWith(Storage value, $Res Function(Storage) then) =
      _$StorageCopyWithImpl<$Res, Storage>;
  @useResult
  $Res call({int id, String name});
}

/// @nodoc
class _$StorageCopyWithImpl<$Res, $Val extends Storage>
    implements $StorageCopyWith<$Res> {
  _$StorageCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? name = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_StorageCopyWith<$Res> implements $StorageCopyWith<$Res> {
  factory _$$_StorageCopyWith(
          _$_Storage value, $Res Function(_$_Storage) then) =
      __$$_StorageCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({int id, String name});
}

/// @nodoc
class __$$_StorageCopyWithImpl<$Res>
    extends _$StorageCopyWithImpl<$Res, _$_Storage>
    implements _$$_StorageCopyWith<$Res> {
  __$$_StorageCopyWithImpl(_$_Storage _value, $Res Function(_$_Storage) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? name = null,
  }) {
    return _then(_$_Storage(
      null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_Storage implements _Storage {
  const _$_Storage(this.id, this.name);

  factory _$_Storage.fromJson(Map<String, dynamic> json) =>
      _$$_StorageFromJson(json);

  @override
  final int id;
  @override
  final String name;

  @override
  String toString() {
    return 'Storage(id: $id, name: $name)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_Storage &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.name, name) || other.name == name));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, id, name);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_StorageCopyWith<_$_Storage> get copyWith =>
      __$$_StorageCopyWithImpl<_$_Storage>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_StorageToJson(
      this,
    );
  }
}

abstract class _Storage implements Storage {
  const factory _Storage(final int id, final String name) = _$_Storage;

  factory _Storage.fromJson(Map<String, dynamic> json) = _$_Storage.fromJson;

  @override
  int get id;
  @override
  String get name;
  @override
  @JsonKey(ignore: true)
  _$$_StorageCopyWith<_$_Storage> get copyWith =>
      throw _privateConstructorUsedError;
}

StorageList _$StorageListFromJson(Map<String, dynamic> json) {
  return _StorageList.fromJson(json);
}

/// @nodoc
mixin _$StorageList {
  List<Storage>? get storageList => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $StorageListCopyWith<StorageList> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $StorageListCopyWith<$Res> {
  factory $StorageListCopyWith(
          StorageList value, $Res Function(StorageList) then) =
      _$StorageListCopyWithImpl<$Res, StorageList>;
  @useResult
  $Res call({List<Storage>? storageList});
}

/// @nodoc
class _$StorageListCopyWithImpl<$Res, $Val extends StorageList>
    implements $StorageListCopyWith<$Res> {
  _$StorageListCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? storageList = freezed,
  }) {
    return _then(_value.copyWith(
      storageList: freezed == storageList
          ? _value.storageList
          : storageList // ignore: cast_nullable_to_non_nullable
              as List<Storage>?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_StorageListCopyWith<$Res>
    implements $StorageListCopyWith<$Res> {
  factory _$$_StorageListCopyWith(
          _$_StorageList value, $Res Function(_$_StorageList) then) =
      __$$_StorageListCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({List<Storage>? storageList});
}

/// @nodoc
class __$$_StorageListCopyWithImpl<$Res>
    extends _$StorageListCopyWithImpl<$Res, _$_StorageList>
    implements _$$_StorageListCopyWith<$Res> {
  __$$_StorageListCopyWithImpl(
      _$_StorageList _value, $Res Function(_$_StorageList) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? storageList = freezed,
  }) {
    return _then(_$_StorageList(
      freezed == storageList
          ? _value._storageList
          : storageList // ignore: cast_nullable_to_non_nullable
              as List<Storage>?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_StorageList implements _StorageList {
  const _$_StorageList(final List<Storage>? storageList)
      : _storageList = storageList;

  factory _$_StorageList.fromJson(Map<String, dynamic> json) =>
      _$$_StorageListFromJson(json);

  final List<Storage>? _storageList;
  @override
  List<Storage>? get storageList {
    final value = _storageList;
    if (value == null) return null;
    if (_storageList is EqualUnmodifiableListView) return _storageList;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(value);
  }

  @override
  String toString() {
    return 'StorageList(storageList: $storageList)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_StorageList &&
            const DeepCollectionEquality()
                .equals(other._storageList, _storageList));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType, const DeepCollectionEquality().hash(_storageList));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_StorageListCopyWith<_$_StorageList> get copyWith =>
      __$$_StorageListCopyWithImpl<_$_StorageList>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_StorageListToJson(
      this,
    );
  }
}

abstract class _StorageList implements StorageList {
  const factory _StorageList(final List<Storage>? storageList) = _$_StorageList;

  factory _StorageList.fromJson(Map<String, dynamic> json) =
      _$_StorageList.fromJson;

  @override
  List<Storage>? get storageList;
  @override
  @JsonKey(ignore: true)
  _$$_StorageListCopyWith<_$_StorageList> get copyWith =>
      throw _privateConstructorUsedError;
}
