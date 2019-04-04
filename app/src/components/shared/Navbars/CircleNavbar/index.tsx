import React, { PureComponent } from 'react';
import CircleNavbarView from './Views';
import { StyleProp, ViewStyle } from 'react-native';

export interface Props {
  onPressBack(): void;
  style?: StyleProp<ViewStyle>;
  circleName: string;
}

class CircleNavbar extends PureComponent<Props> {
  render() {
    return <CircleNavbarView {...this.props} />;
  }
}

export default CircleNavbar;
