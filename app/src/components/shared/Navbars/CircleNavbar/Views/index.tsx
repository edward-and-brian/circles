import React, { PureComponent } from 'react';
import { NavigationScreenProp } from 'react-navigation';
import {
  View,
  Text,
  TouchableOpacity,
  Image,
  ViewStyle,
  StyleProp,
} from 'react-native';
import CircleAvatarButton from '../../../Avatar/CircleAvatarButton';
import { Images } from '../../../../../themes';
import styles from './styles';

export interface Props {
  onPressBack(): void;
  style?: StyleProp<ViewStyle>;
}

class CircleNavbarView extends PureComponent<Props> {
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
          <CircleAvatarButton />
        </View>
      </View>
    );
  }
}

export default CircleNavbarView;
