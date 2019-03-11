import React, { PureComponent } from 'react';
import { View, TouchableOpacity } from 'react-native';
import styles from './styles';

class CircleAvatarButtonView extends PureComponent {
  render() {
    return (
      <TouchableOpacity style={styles.container}>
        <View style={styles.avatar} />
      </TouchableOpacity>
    );
  }
}

export default CircleAvatarButtonView;
