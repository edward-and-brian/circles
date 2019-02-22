import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import CircleNavbar from '../../../components/Base/Navbars/CircleNavbar';
import styles from './styles';

class CircleScreenView extends PureComponent {
  render() {
    return (
      <View style={styles.container}>
        <CircleNavbar />
        <View style={styles.conversationContainer} />
        <View style={styles.footerContainer} />
      </View>
    );
  }
}

export default CircleScreenView;
