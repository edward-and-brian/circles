import React, { PureComponent } from 'react';
import { ScrollView, View, Text, Animated } from 'react-native';
import CircleMessageWindow from '../../../components/circle/MessageWindow';
import CircleNavbar from '../../../components/shared/Navbars/CircleNavbar/';
import CircleFooter from '../../../components/circle/Footer';
import styles from './styles';

export interface Props {
  onPressBack(): void;
}

class CircleScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <CircleNavbar
          onPressBack={this.props.onPressBack}
          style={styles.headerContainer}
        />
        <CircleMessageWindow />
        <CircleFooter />
      </View>
    );
  }
}

export default CircleScreenView;
