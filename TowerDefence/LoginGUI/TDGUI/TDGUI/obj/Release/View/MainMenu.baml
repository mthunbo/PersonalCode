   M S B A M L   `   `   `  ���� @  <TDGUI, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null'clr-namespace:TDGUI.View
TDGUI.View      TDGUI.View.MainMenu   R NWindowsBase, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35\ XPresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35W SPresentationCore, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35D 9http://schemas.microsoft.com/winfx/2006/xaml/presentation    R NSystem.Xaml, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089:x,http://schemas.microsoft.com/winfx/2006/xaml     @d2http://schemas.microsoft.com/expression/blend/2008     Jmc;http://schemas.openxmlformats.org/markup-compatibility/2006     *localclr-namespace:TDGUI.View       � Title$  MainMenu��$	��600��$	��400��
���   �$	��400��$	��400��$	��600��$	��600�� � DataContext ��    Locator �� Source# [   �� Path$ 
MainMenuVM�.���  � RowDefinitions �� $��1* ��� $��1* ��� $��1* ��� $��1* �.+��� $"��../Resources/ArcherTower.png��$��93,-2,157.333,28.667q� J� OpacityMask
 �   �z���  �� ScaleX$ -1Q��� $"��../Resources/ArcherTower.png��$��298,0,-29.667,26.667q�
 �   ��� $"��../Resources/CannonTower.png��$��96,29,177.333,-2.333q�
 �   �$��3��z�V� .~��� $ -1Q��� $"��../Resources/CannonTower.png��$��298,27,-24.667,-0.333q�
 �   �$��3���� -    ��Name$ 	PlayTDbtn��$��0��$��Center=�$��Center=�
������ $��0q�
������ ���� * StartCommandApplicationCommand+����  �� CornerRadius$ 5��$��34��$	��200��$��0q�
�������$��2,2,2,2q�._��� $��Play Tower defence��$��Center=�$��Center=�
��������� -   $ seePrHighbtn��$��1��$��Center=�$��Center=�
������ $��0q�
������ ���� *CreateHighscoreView+���� $ 5��$��34��$	��200��$��0q�
�������$��2,2,2,2q�._��� $��See Personal Highscore��$��Center=�$��Center=�
��������� -   $ seeHighscbtn��$��2��$��Center=�$��Center=�
������ $��0q�
������ ���� *CreateGuestHighscoreView+���� $ 5��$��34��$	��200��$��0q�
�������$��2,2,2,2q�._��� $��See Highscore��$��Center=�$��Center=�
��������� -   $ 	logoutbtn��$��3��$��Center=�$��Center=�
������ $��0q�
������ ���� *LogoutCommand+���� $ 5��$��34��$	��200��$��0q�
�������$��2,2,2,2q�._��� $��Log out��$��Center=�$��Center=�
�������