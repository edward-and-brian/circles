import React, { PureComponent } from 'react';
import { Colors } from '../../../../themes';
import LoginButtonView from './Views';

export interface Props {
  text: string;
  color: string;
  fontColor: string;
  onPress(): void;
}

class LoginButton extends PureComponent<Props> {
  static defaultProps: Partial<Props> = {
    fontColor: Colors.black,
  };

  render() {
    return <LoginButtonView {...this.props} />;
  }
}

export default LoginButton;
