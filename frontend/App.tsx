import { StatusBar } from 'expo-status-bar';
import React, {useState, useEffect} from 'react';
import { Alert, Button,TouchableOpacity, StyleSheet, Text, View } from 'react-native';

const XpRender = () => {
  const [isXp, setIsXp] = useState(0);
  if (isXp >= 100) {
    setIsXp(0);
  }
    const onPressButton = () => {
      setIsXp(isXp + 10);
    };
  return (
    <View style={styles.container}>
      <Text style={{color: 'orange'}}> 
        Hello I'm an XP bar
      </Text>
      <TouchableOpacity onPress={onPressButton} style={styles.button}>
        <Text style={styles.buttonText}> 10 XP </Text>
      </TouchableOpacity>
      <Text style= {{color:'green'}}>
        {isXp}
      </Text>
    </View>
  )
}

export default function App() {
  return (
    <XpRender></XpRender>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#000',
    alignItems: 'center',
    justifyContent: 'center',
  },
  button: {
    padding:10,
    width: 100,
    backgroundColor: '#b703ff',
    borderRadius: 4,
    alignItems: 'center'
  },
  buttonText: {
    color:'white',
  }
});
