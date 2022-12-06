import 'package:freezed_annotation/freezed_annotation.dart';

part 'notification.freezed.dart';
part 'notification.g.dart';

@freezed
abstract class Notification with _$Notification {
  const factory Notification(
    String title,
    String message,
    @JsonKey(name: 'created_date') String createdDate,
  ) = _Notification;
  factory Notification.fromJson(Map<String, dynamic> json) =>
      _$NotificationFromJson(json);
}

@freezed
abstract class NotificationList with _$NotificationList {
  const factory NotificationList(List<Notification>? messageList) =
      _NotificationList;
  factory NotificationList.fromJson(Map<String, dynamic> json) =>
      _$NotificationListFromJson(json);
}
