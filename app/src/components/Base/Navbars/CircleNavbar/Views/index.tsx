import React, { PureComponent } from 'react';
import { NavigationScreenProp } from 'react-navigation';
import { View, Text, TouchableOpacity, Image } from 'react-native';
import { Images } from '../../../../../themes';
import styles from './styles';

export interface Props {
  onPressBack(): void;
}

class CircleNavbarView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
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
        <View style={styles.avatarContainer} />
      </View>
    );
  }
}

export default CircleNavbarView;
