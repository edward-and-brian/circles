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
  inputsContainer: {
    flex: 2,
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  buttonsContainer: {
    flex: 3,
    alignItems: 'center',
  },
  buttonContainer: {
    height: Scaled.screen.height / 20,
    width: '45%',
    marginBottom: Scaled.screen.height / 100
  },
});
