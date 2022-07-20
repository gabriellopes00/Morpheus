import { MigrationInterface, QueryRunner, Table, TableForeignKey } from 'typeorm'

export class CreateFavoritesMigration implements MigrationInterface {
  name = this.constructor.name.concat('1658315610357')

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'favorites',
        columns: [
          { name: 'id', type: 'uuid', isPrimary: true },
          { name: 'account_id', type: 'uuid', isNullable: false },
          { name: 'event_id', type: 'uuid', isNullable: false }
        ]
      }),
      true
    )
    await queryRunner.createForeignKey(
      'favorites',
      new TableForeignKey({
        name: 'fk_accounts_favorites',
        referencedTableName: 'accounts',
        referencedColumnNames: ['id'],
        columnNames: ['account_id'],
        onDelete: 'cascade',
        onUpdate: 'cascade'
      })
    )
    await queryRunner.createForeignKey(
      'favorites',
      new TableForeignKey({
        name: 'fk_events_favorites',
        referencedTableName: 'events',
        referencedColumnNames: ['id'],
        columnNames: ['event_id'],
        onDelete: 'cascade',
        onUpdate: 'cascade'
      })
    )
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropForeignKey('favorites', 'fk_events_favorites')
    await queryRunner.dropForeignKey('favorites', 'fk_accounts_favorites')
    await queryRunner.dropTable('favorites', true)
  }
}
