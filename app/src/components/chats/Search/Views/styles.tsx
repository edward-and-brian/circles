import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../themes';

export default StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
  },
  searchContainer: {
    flex: 7,
    justifyContent: 'center',
    alignItems: 'center',
  },
  messageInput: {
    width: '95%',
    justifyContent: 'center',
    height: Scaled.screen.height * 0.045,
    paddingLeft: Scaled.screen.height * 0.013,
    borderRadius: Scaled.screen.height * 0.008,
    fontFamily: Fonts.book,
    fontSize: Scaled.fontSize.h10,
    backgroundColor: Colors.lightGray,
  },
  buttonContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  buttonImage: {
    width: '50%',
    height: '50%',
  },
});
