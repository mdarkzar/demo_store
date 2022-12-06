// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'notification.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_Notification _$$_NotificationFromJson(Map<String, dynamic> json) =>
    _$_Notification(
      json['title'] as String,
      json['message'] as String,
      json['created_date'] as String,
    );

Map<String, dynamic> _$$_NotificationToJson(_$_Notification instance) =>
    <String, dynamic>{
      'title': instance.title,
      'message': instance.message,
      'created_date': instance.createdDate,
    };

_$_NotificationList _$$_NotificationListFromJson(Map<String, dynamic> json) =>
    _$_NotificationList(
      (json['messageList'] as List<dynamic>?)
          ?.map((e) => Notification.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$$_NotificationListToJson(_$_NotificationList instance) =>
    <String, dynamic>{
      'messageList': instance.messageList,
    };
