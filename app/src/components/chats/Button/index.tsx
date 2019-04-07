import React, { PureComponent } from 'react';
import ButtonView from './Views';

export interface Props {
  chat: number;
  onPressChat(): void;
}

class Button extends PureComponent<Props> {
  render() {
    return <ButtonView {...this.props} />;
  }
}

export default Button;
