import 'dart:convert';
import 'package:hex/hex.dart';
import 'package:crypto/crypto.dart';

String toJson(dynamic object) {
  var encoder = JsonEncoder.withIndent('     ');
  return encoder.convert(object);
}

String hashStringSHA256(String input) {
  var bytes = utf8.encode(input);
  var digest = sha256.convert(bytes);
  return toHexString(digest.bytes);
}

String toHexString(List<int> data) {
  return HEX.encode(data);
}
