singularity = {};

singularity.board = function() {
  "use strict";

  /*jshint browser:true, latedef:false */

  var color = [
    'rgba(252,159,255, {{opacity}})',
    'rgba(250,111,207, {{opacity}})',
    'rgba(201,157,220, {{opacity}})'
  ];

  color.opacity = "0.3";

  color.random = function() {
    var c = color[~~(Math.random() * 10 % color.length)];
    c = c.replace('{{opacity}}', this.opacity);
    return c;
  };

  color.border = function() {
    var c = color[~~(Math.random() * 10 % color.length)];
    c = c.replace('{{opacity}}', "0.15");
    return c;
  };

  var config = {
    height: 800,
    width: 800,
    padding: {
      x: 90,
      y: 50
    }
  };

  var players = {},
      layers = {};

  function createCanvas(id) {
    var canvas = document.createElement('canvas'),
        context = canvas.getContext('2d');

    canvas.id = id;
    canvas.setAttribute('width', config.width + config.padding.x + 1);
    canvas.setAttribute('height', config.height + config.padding.y + 1);

    canvas.style.position = 'absolute';

    var board = document.getElementById('board');
    board.appendChild(canvas);

    return {
      canvas: canvas,
      context: context
    };
  }

  return {
    drawBackground: function() {

      var context = layers.background.context,
          canvas = layers.background.canvas,
          x,y;

      canvas.width = canvas.width;

      for( x = 0.5; x <= canvas.width + 0.5; x = x + 10 ) {
        for( y = 0.5; y <= canvas.height + 0.5; y = y + 10 ) {

          if( x <= 10.5 || y <= 10.5 ||
              x >= canvas.width - 20.5 || 
              y >= canvas.height - 20.5 ) {
            context.fillStyle = color.border();
          } else {
            context.fillStyle = color.random();
          }

          context.clearRect(x, y, 10, 10);
          context.fillRect(x, y, 10, 10);
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
        layer.context.drawImage(
          player.avatar, 
          position.x - player.avatar.width /2, 
          position.y - player.avatar.height /2
        );

        var context = layer.context;

        context.font = '10px pirulen';
        context.textBaseline = 'top';
        context.fillStyle = 'white';

        var x = position.x - player.avatar.width /2,
            y = position.y - player.avatar.height /2;
        
        x = x - (7 * (player.name.length /2));
        y = y - 20;
        context.fillText(player.name, x, y);
      }
    },

    set: function(c) {
      if (typeof c !== 'undefined') {
        config = c;
      }
    },

    start: function() {
      layers.background = createCanvas('background');
      layers.game = createCanvas('game');

      this.drawBackground();
      setInterval(this.drawBackground, 3000);
      setInterval(this.drawGame, 100);
    },

    update: function(player) {
      // transform to cartisian coordinates
      player.position.y = config.height - player.position.y;

      player.position.x += config.padding.x /2;
      player.position.y += config.padding.y /2;

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
    },

    remove: function(player) {
      delete players[player.name];
    }
  };

};
