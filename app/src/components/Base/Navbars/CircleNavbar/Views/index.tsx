import React, { PureComponent } from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles';

class CircleScreenView extends PureComponent {
  render() {
    return (
      <View style={styles.container}>
        <TouchableOpacity>
          <View style={styles.backArrow}/>
        </TouchableOpacity>
        <View style={styles.titleContainer}>
          <Text style={styles.chatTitle} />
          <Text style={styles.circleTitle} />
        </View>
      </View>
    );
  }
}

export default CircleScreenView;
