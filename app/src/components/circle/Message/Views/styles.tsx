import { StyleSheet } from 'react-native';
import { Colors, Fonts, Scaled } from '../../../../themes/';

export default StyleSheet.create({
  container: {
    flexDirection: 'row',
    marginBottom: Scaled.screen.width * 0.02,
    marginRight: Scaled.screen.width * 0.04,
  },
  avatarContainer: {
    marginHorizontal: Scaled.screen.width * 0.015,
  },
  avatar: {},
  contentContainer: {
    flex: 6,
  },
  nameAndDateContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    height: Scaled.screen.width * 0.06,
  },
  name: {
    fontFamily: Fonts.heavy,
    fontSize: Scaled.fontSize.h10,
    marginRight: Scaled.screen.width * 0.02,
  },
  date: {
    fontFamily: Fonts.book,
    fontSize: Scaled.fontSize.h11,
    color: Colors.gray,
  },
  message: {
    fontFamily: Fonts.book,
    fontSize: Scaled.fontSize.h10,
    marginBottom: Scaled.screen.width * 0.01,
    lineHeight:  Scaled.screen.width * 0.044
  },
});
