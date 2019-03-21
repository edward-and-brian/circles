import React, { PureComponent } from 'react';
import {
  ScrollView,
  Text,
  SectionList,
  View,
  SectionListData,
} from 'react-native';
import CircleMessageDivider from '../../MessageDivider';
import CircleMessage from '../../Message';
import styles from './styles';

interface Message {
  user: {
    name: string;
    avatar: string;
  };
  messages: string[];
  date: Date;
}

interface Conversation {
  title: string;
  data: Message[];
}

export interface Props {
  conversation: Conversation[];
}

const renderItem = ({ item }: { item: Message }) => {
  return <CircleMessage messageGroup={item} />;
};

const renderSectionHeader = ({
  section: { title },
}: {
  section: SectionListData<Conversation>;
}) => {
  return <CircleMessageDivider label={title} />;
};

class MessageWindowView extends PureComponent<Props> {
  render() {
    return (
      <SectionList
        renderItem={renderItem}
        renderSectionFooter={renderSectionHeader}
        sections={this.props.conversation}
        keyExtractor={(item, index) => item + index}
        style={styles.container}
        inverted
      />
    );
  }
}

export default MessageWindowView;
