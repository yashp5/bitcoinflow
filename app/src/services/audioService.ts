class AudioService {
  constructor() {
    this.audioContext = new (window.AudioContext ||
      window.webkitAudioContext)();
  }

  handleOption(option) {
    let frequency = 440; // Default frequency
    let type = "sine"; // Default waveform type

    // Define audio characteristics for each scenario
    if (option.dir === "BUY" && option.cp === "CALL") {
      // Bullish scenario
      frequency = 880; // Higher frequency for bullish
      type = "square"; // Distinct waveform
    } else if (option.dir === "SELL" && option.cp === "PUT") {
      // Bullish scenario
      frequency = 880; // Same higher frequency for bullish
      type = "triangle"; // Different waveform for variation
    } else if (option.dir === "BUY" && option.cp === "PUT") {
      // Bearish scenario
      frequency = 440; // Lower frequency for bearish
      type = "sawtooth"; // Distinct waveform
    } else if (option.dir === "SELL" && option.cp === "CALL") {
      // Bearish scenario
      frequency = 440; // Same lower frequency for bearish
      type = "sine"; // Different waveform for variation
    }

    this.playSound(frequency, 500, type);
  }

  playSound(frequency = 440, duration = 200, type = "sine") {
    // Reduced duration to 200ms
    if (this.audioContext.state === "suspended") {
      this.audioContext.resume();
    }

    const oscillator = this.audioContext.createOscillator();
    const gainNode = this.audioContext.createGain();

    oscillator.type = type;
    oscillator.frequency.setValueAtTime(
      frequency,
      this.audioContext.currentTime
    );
    gainNode.gain.setValueAtTime(0.1, this.audioContext.currentTime);

    oscillator.connect(gainNode);
    gainNode.connect(this.audioContext.destination);

    oscillator.start();
    oscillator.stop(this.audioContext.currentTime + duration / 1000); // Adjusted duration
  }
}

const audioService = new AudioService();
export default audioService;
