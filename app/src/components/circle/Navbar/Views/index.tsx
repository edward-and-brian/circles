import React, { PureComponent } from 'react';
import {
  View,
  Text,
  TouchableOpacity,
  Image,
  ViewStyle,
  StyleProp,
} from 'react-native';
import AvatarButton from '../../../shared/Button/AvatarButton';
import { Images, Scaled } from '../../../../themes';
import styles from './styles';

export interface Props {
  onPressBack(): void;
  style?: StyleProp<ViewStyle>;
}

class NavbarView extends PureComponent<Props> {
  render() {
    return (
      <View style={[styles.container, this.props.style]}>
        <TouchableOpacity
          style={styles.arrowContainer}
          onPress={this.props.onPressBack}
        >
          <Image
            source={Images.backArrow}
            style={styles.backArrow}
            resizeMode="contain"
          />
        </TouchableOpacity>
        <View style={styles.titleContainer}>
          <Text style={styles.chatTitle}>The Come Back Kids</Text>
          <Text style={styles.circleTitle}>Kill the Birds</Text>
        </View>
        <View style={styles.avatarContainer}>
          <AvatarButton diameter={Scaled.screen.height * 0.055} clickable />
        </View>
      </View>
    );
  }
}

export default NavbarView;
