import 'package:flutter/widgets.dart';
import 'package:sizer/sizer.dart';

Widget pB(Widget child, double bottom) {
  return Padding(child: child, padding: EdgeInsets.only(bottom: bottom));
}

Widget centerSingleChildScroll({required Widget child, double? width}) {
  return Center(
      child:
          SingleChildScrollView(child: SizedBox(width: width, child: child)));
}

Widget wSpace(double space) {
  return SizedBox(width: space.w);
}

Widget hSpace(double space) {
  return SizedBox(height: space.h);
}
