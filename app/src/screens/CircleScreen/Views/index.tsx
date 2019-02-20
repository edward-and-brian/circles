import React, { PureComponent } from 'react';
import { View } from 'react-native';
import CircleNavbar from '../../../components/Base/Navbars/CircleNavbar';
import styles from './styles';

class CircleScreenView extends PureComponent {
  render() {
    return (
      <View style={styles.container}>
        <CircleNavbar />
      </View>
    );
  }
}

export default CircleScreenView;
