import { StyleSheet } from 'react-native';
import { Colors, Fonts, Scaled } from '../../../themes/';

export default StyleSheet.create({
  container: {
    flex: 1,
  },
  titleContainer: {
    flex: 2,
    alignItems: 'center',
    justifyContent: 'flex-end',
  },
  title: {
    fontFamily: Fonts.medium,
    fontSize: Scaled.fontSize.h1,
  },
  inputContainer: {
    flex: 2,
    backgroundColor: Colors.sunset1,
  },
  buttonContainer: {
    flex: 3,
    backgroundColor: Colors.sunset2,
  },
});
