   M S B A M L   `   `   `  ���� @  <TDGUI, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null'clr-namespace:TDGUI.View
TDGUI.View  -clr-namespace:CustomControlCustomControl  !    TDGUI.View.RegisterUserView   R NWindowsBase, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35\ XPresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35W SPresentationCore, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35D 9http://schemas.microsoft.com/winfx/2006/xaml/presentation    R NSystem.Xaml, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089:x,http://schemas.microsoft.com/winfx/2006/xaml     @d2http://schemas.microsoft.com/expression/blend/2008     Jmc;http://schemas.openxmlformats.org/markup-compatibility/2006     *localclr-namespace:TDGUI.View     /controlclr-namespace:CustomControl     
���   �  � Title$  Create User��$	��600��$	��400��$	��600��$	��600��$	��400��$	��400�� � DataContext ��    Locator �� Source# [   �� Path$ CreateUserVM�.��� ����  �� ImageSource$  /Resources/CannonTower.png�� � RowDefinitions �� $
��1.5* ��� $
��1.3* ��� $
��1.5* �.+��� $2��,Enter username and password to be registered��$��0��$��Center=�$	��199��$��0,0,0.333,0.333q�$��Bottom=�
�������$��17(��� $��	Username:��$��1��$��50,39.667,0,112q�$
��Left=�$��68��
��������� -    ��Name$ Username��$	��150��$��25��$��169,23,74.333,98.527q�$��1��$��Center=�
����������� *
Username+�� $��	Password:��$��1��$��50,91.667,0,60q�$
��Left=�$��68��
�������'   !CustomControl.BindablePasswordBox  -     Name$ Password��$	��150��$��25��$��169,75,74.333,47.333q�$��1��$��Center=�   Password �� *
Password+�� -   	 ��Name$	 BackToLoginBTN��$��2��$��200,20,0,100q�$��40��$	��110��
������ $��0q�
������ ���� *backToLogin+���� 
 �� CornerRadius$
 5��$��34��$	��108��$��0q�
�������$��2,2,2,2q�
���   �._��� $��Back to login��$��Center=�$��Center=�
��������� -   $	 registerUsrBTN��$��2��$��20,20,250,100q�$��40��$	��110��
������ $��0q�
������ ���� *CreateUserCommand+���� $
 5��$��34��$	��108��$��0q�
�������$��2,2,2,2q�
���   �._��� $��Register��$��Center=�$��Center=�
�������