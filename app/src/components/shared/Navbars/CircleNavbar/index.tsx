import React, { PureComponent } from 'react';
import CircleNavbarView from './Views';
import { StyleProp, ViewStyle } from 'react-native';

export interface Props {
  onPressBack(): void;
  style?: StyleProp<ViewStyle>;
}

class CircleNavbar extends PureComponent<Props> {
  render() {
    return <CircleNavbarView {...this.props} />;
  }
}

export default CircleNavbar;
