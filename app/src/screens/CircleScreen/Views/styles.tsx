import { StyleSheet } from 'react-native';
import { Colors, Fonts, Scaled } from '../../../themes/';

export default StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.white,
  },
  conversationContainer: {
    flex: 1,
  },
  headerContainer: {
    height: Scaled.navBarHeight,
  },
  light: {
    fontSize: 25,
    fontFamily: Fonts.light,
  },
  book: {
    fontSize: 25,
    fontFamily: Fonts.book,
  },
  medium: {
    fontSize: 25,
    fontFamily: Fonts.medium,
  },
  heavy: {
    fontSize: 25,
    fontFamily: Fonts.heavy,
  },
  black: {
    fontSize: 25,
    fontFamily: Fonts.black,
  },
});
