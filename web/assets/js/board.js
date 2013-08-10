singularity = {};

singularity.board = function() {
  "use strict";

  /*jshint browser:true, latedef:false */

  var color = [
    'rgba(252,159,255, 0.3)',
    'rgba(250,111,207, 0.3)',
    'rgba(201,157,220, 0.3)'
  ];

  color.random = function() {
    return color[~~(Math.random() * 10 % color.length)];
  };

  var config = {
    height: 800,
    width: 800
  };

  var players = {},
      layers = {};

  function createCanvas(id) {
    var canvas = document.createElement('canvas'),
        context = canvas.getContext('2d');

    canvas.id = id;
    canvas.setAttribute('width', config.width + 1);
    canvas.setAttribute('height', config.height + 1);

    canvas.style.margin = '0 auto';
    canvas.style.position = 'absolute';

    document.body.appendChild(canvas);

    return {
      canvas: canvas,
      context: context
    };
  }

  return {
    drawBackground: function() {

      var context = layers.background.context;
      context.fillStyle = color.random();

      for( var x = 0.5; x <= config.width + 0.5; x = x + 10 ) {
        for( var y = 0.5; y <= config.height + 0.5; y = y + 10 ) {
          context.clearRect(x, y, 10, 10);
          context.fillRect(x, y, 10, 10);
          context.fillStyle = color.random();
        }
      }
    },

    drawGame: function() {
      var layer = layers.game;

      // force a reset by touching the canvas 
      // size.
      layer.canvas.width = layer.canvas.width;
      
      for( var name in players ) {
        var player = players[name];
        var position = player.position;
        layer.context.drawImage(player.avatar, position.x, position.y);
      }
    },

    set: function(c) {
      config = c;
    },

    start: function() {
      layers.background = createCanvas('background');
      layers.game = createCanvas('game');

      this.drawBackground();
      setInterval(this.drawBackground, 3000);
      setInterval(this.drawGame, 100);
    },

    update: function(player) {
      if (typeof players[player.name] === 'undefined') {
        var img = new Image();
        img.src = 'images/player.png';
        img.onload = function() {
          player.avatar = img;
          players[player.name] = player;
          layers.game.context.drawImage(img, 0, 0);
        };
        return;
      }

      players[player.name].position = player.position;
    }
  };

};
