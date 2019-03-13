import { StyleSheet } from 'react-native';
import { Colors, Fonts, Scaled } from '../../../themes/';

export default StyleSheet.create({
  container: {
    height: Scaled.screen.width * 0.06,
    flexDirection: 'row',
    alignItems: 'center',
    // backgroundColor: 'lightpink'
  },
  leftSpacer: {
    flex: 6,
    height: 0.6,
    backgroundColor: Colors.gray,
  },
  label: {
    paddingHorizontal: Scaled.screen.width * 0.02,
    fontSize: Scaled.fontSize.h10,
    fontFamily: Fonts.medium,
    color: Colors.gray,
  },
  rightSpacer: {
    flex: 1,
    height: 0.6,
    backgroundColor: Colors.gray,
  },
});
