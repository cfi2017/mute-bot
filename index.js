const Discord = require('discord.js')
const client = new Discord.Client();
const process = require('process');

client.on('ready', () => {
    console.log(`Logged in as ${client.user.tag}!`);
});

let muted = false;

client.on('message', msg => {
    if (msg.content === '!mute') {
        if (muted) {
            muted = false;
            msg.reply('unmuting users.');
            msg.guild.members.fetch().then(members => {
                for (let [id, member] of members) {
                    member.edit({mute: false}).then(() => console.log('done'));
                }
            })
        } else {
            muted = true;
            msg.reply('muting users.');
            msg.guild.members.fetch().then(members => {
                for (let [id, member] of members) {
                    member.edit({mute: true}).then(() => console.log('done'));
                }
            })
        }
    }
});

client.login('NzQzMjM2ODM3MDkwMzk0MTU2.XzRvPg.MIu1nJa38pjFnIDSbItngU-J1BI');