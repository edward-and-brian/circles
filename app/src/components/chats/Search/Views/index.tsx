import React, { PureComponent } from 'react';
import { View, ScrollView, TouchableOpacity } from 'react-native';
import styles from './styles';

export interface Props {}

class ScrollListView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.searchContainer} />
        <TouchableOpacity style={styles.buttonContainer} />
      </View>
    );
  }
}

export default ScrollListView;
