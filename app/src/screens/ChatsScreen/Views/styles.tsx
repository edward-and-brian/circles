import { StyleSheet } from 'react-native';
import { Colors, Fonts, Scaled } from '../../../themes/';

export default StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.white,
  },
  searchContainer: {
    flex: 1,
    backgroundColor: Colors.green,
  },
  chatsContainer: {
    flex: 10,
    backgroundColor: Colors.yellow,
  },
  chatContainer: {
    height: Scaled.screen.height * 0.1,
    backgroundColor: Colors.white,
    margin: Scaled.screen.height * 0.006,
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: Scaled.screen.height * 0.01,
  },
  chatName: {
    fontSize: Scaled.fontSize.h4,
    fontFamily: Fonts.heavy,
  },
  spacer: {
    height: Scaled.screen.height * 0.03,
  },
});
