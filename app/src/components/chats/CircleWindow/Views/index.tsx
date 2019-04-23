import React, { PureComponent } from 'react';
import { Animated } from 'react-native';
import moment from 'moment';
import { Circle } from '../../ScrollList';
import CircleButton from '../../CircleButton';

import styles from './styles';

export interface Props {
  circles: Circle[];
  circleWindowHeight: Animated.Value;
  onPressCircle(): void;
}

const date = moment().subtract(4, 'hour');

class CircleWindowView extends PureComponent<Props> {
  render() {
    return (
      <Animated.ScrollView
        style={[
          styles.circleContainer,
          { maxHeight: this.props.circleWindowHeight },
        ]}
      >
        {this.props.circles.map(circle => (
          <CircleButton
            circle={circle}
            onPressCircle={this.props.onPressCircle}
            key={`circle${circle.name}`}
          />
        ))}
      </Animated.ScrollView>
    );
  }
}

export default CircleWindowView;
