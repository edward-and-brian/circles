import React, { PureComponent } from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles';

class CircleNavbarView extends PureComponent {
  render() {
    return (
      <View style={styles.container}>
          <TouchableOpacity style={styles.arrowContainer}>
            <View style={styles.backArrow} />
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
