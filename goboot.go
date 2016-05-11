package main

import (
    "syscall"
)

/*
        syscall commands
        LINUX_REBOOT_CMD_CAD_OFF         = 0x0
        LINUX_REBOOT_CMD_CAD_ON          = 0x89abcdef
        LINUX_REBOOT_CMD_HALT            = 0xcdef0123
        LINUX_REBOOT_CMD_KEXEC           = 0x45584543
        LINUX_REBOOT_CMD_POWER_OFF       = 0x4321fedc
        LINUX_REBOOT_CMD_RESTART         = 0x1234567
        LINUX_REBOOT_CMD_RESTART2        = 0xa1b2c3d4
        LINUX_REBOOT_CMD_SW_SUSPEND      = 0xd000fce2
        LINUX_REBOOT_MAGIC1              = 0xfee1dead
        LINUX_REBOOT_MAGIC2              = 0x28121969

       Linux Man page info
       LINUX_REBOOT_CMD_CAD_OFF
              (RB_DISABLE_CAD, 0).  CAD is disabled.  This means that the
              CAD keystroke will cause a SIGINT signal to be sent to init
              (process 1), whereupon this process may decide upon a proper
              action (maybe: kill all processes, sync, reboot).

       LINUX_REBOOT_CMD_CAD_ON
              (RB_ENABLE_CAD, 0x89abcdef).  CAD is enabled.  This means that
              the CAD keystroke will immediately cause the action associated
              with LINUX_REBOOT_CMD_RESTART.

       LINUX_REBOOT_CMD_HALT
              (RB_HALT_SYSTEM, 0xcdef0123; since Linux 1.1.76).  The message
              "System halted." is printed, and the system is halted.
              Control is given to the ROM monitor, if there is one.  If not
              preceded by a sync(2), data will be lost.

       LINUX_REBOOT_CMD_KEXEC
              (RB_KEXEC, 0x45584543, since Linux 2.6.13).  Execute a kernel
              that has been loaded earlier with kexec_load(2).  This option
              is available only if the kernel was configured with
              CONFIG_KEXEC.

       LINUX_REBOOT_CMD_POWER_OFF
              (RB_POWER_OFF, 0x4321fedc; since Linux 2.1.30).  The message
              "Power down." is printed, the system is stopped, and all power
              is removed from the system, if possible.  If not preceded by a
              sync(2), data will be lost.

       LINUX_REBOOT_CMD_RESTART
              (RB_AUTOBOOT, 0x1234567).  The message "Restarting system." is
              printed, and a default restart is performed immediately.  If
              not preceded by a sync(2), data will be lost.

       LINUX_REBOOT_CMD_RESTART2
              (0xa1b2c3d4; since Linux 2.1.30).  The message "Restarting
              system with command '%s'" is printed, and a restart (using the
              command string given in arg) is performed immediately.  If not
              preceded by a sync(2), data will be lost.

       LINUX_REBOOT_CMD_SW_SUSPEND
              (RB_SW_SUSPEND, 0xd000fce1; since Linux 2.5.18).  The system
              is suspended (hibernated) to disk.  This option is available
              only if the kernel was configured with CONFIG_HIBERNATION.

*/

func main() {

    syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)

}
