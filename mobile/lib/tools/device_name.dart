import 'package:device_info_plus/device_info_plus.dart';
import 'package:universal_platform/universal_platform.dart';

Future<String> deviceName() async {
  final DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();

  try {
    if (UniversalPlatform.isAndroid) {
      final androidInfo = await deviceInfo.androidInfo;
      return androidInfo.model ?? 'android device';
    } else if (UniversalPlatform.isLinux) {
      final linuxDeviceInfo = await deviceInfo.linuxInfo;
      return linuxDeviceInfo.prettyName;
    } else if (UniversalPlatform.isMacOS) {
      final macOsDeviceInfo = await deviceInfo.macOsInfo;
      return macOsDeviceInfo.model;
    } else if (UniversalPlatform.isWeb) {
      final webBrowserInfo = await deviceInfo.webBrowserInfo;
      return webBrowserInfo.userAgent ?? 'web device';
    } else {
      return "unknown device";
    }
  } catch (err) {
    return "unknown device";
  }
}
